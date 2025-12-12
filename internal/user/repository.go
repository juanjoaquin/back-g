package user

import (
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Create(user *User) error      // Le pasamos como puntero al User
	GetAll() ([]User, error)      // El Get all, nos devuelve un array de usuarios
	Get(id string) (*User, error) // El Get by ID, nos devuelve un ID, y un puntero de User
	Delete(id string) error
	Update(id string, firstName *string, lastName *string, email *string, phone *string) error
}

// Esta struct va hacer referencia a la DB de GORM
type repo struct {
	log *log.Logger
	db  *gorm.DB
}

// Creamos una funcion que va a instanciar este repo.

func NewRepo(log *log.Logger, db *gorm.DB) Repository {
	return &repo{
		log: log,
		db:  db,
	}
}

// Creamos el Metodo Create
func (repo *repo) Create(user *User) error {
	repo.log.Println(user) // Este loguer es el que nosotros le hemos pasado. Es lo que imprimira al pegarle al POST

	/* Nuestro ID es UUID, entonces debemos definirle ese UUID desde esta capa (tenemos que usar el package de google/uuid) */
	user.ID = uuid.New().String()

	/* Tenemos que hacer del objeto  de "db" el metodo "Create", llamando a nuestra Struct (repo) que le debemos pasar la entidad del User */
	result := repo.db.Create(user)

	// Tenemos 2 tipos de manejos de error. Este en el que le decimos, que si el resultado da error, y es distinto a null que lo tire:

	if result.Error != nil {
		repo.log.Println("[ERROR]-[REPOSITORY]-[CREATE]", result.Error)
		return result.Error
	}

	// O este donde seteamos con la funcion propia en la creacion del User, y no una vez previamente declara como en la primera opcion
	/* 	if err := repo.db.Create(user).Error; err != nil {
		repo.log.Println(err)
		return err
	} */

	repo.log.Println("User creado exitosamente", user.ID)

	return nil
}

// Creamo el Metodo Get All
func (repo *repo) GetAll() ([]User, error) {
	var u []User // Declaramos la variable user. Que sera un vector de usuarios

	/* Utilizamos la funcion de nuestro repo, para tener la DB, y ejecutar el metodo "Model"
	Con esto especificamos el Modelo que vamos a utilizar. En este caso el User, con su puntero */
	result := repo.db.Model(&u).Order("created_at desc").Find(&u) // Le aplicamos un orderBy, y un Find para encontrar el user

	// Hanldeamos el error
	if result.Error != nil {
		return nil, result.Error
	}

	// Returnamos el user y el nil
	return u, nil

}

// Creamo el Metodo Get By ID
func (repo *repo) Get(id string) (*User, error) {
	/* Primero debemos generar una estructura User para poder pasarle el ID a GORM */
	user := User{ID: id}

	/* Para buscar la informacion, utilizamos el .First() con el puntero en el User.  */
	result := repo.db.First(&user) // First es el primer elemento que encuentra

	// Handleamos el error
	if result.Error != nil {
		return nil, result.Error
	}

	// Devolvemos al puntero del User, tanto como el nil. No se devuelve el result
	return &user, nil

}

// Creamos el Metodo DELETE
func (repo *repo) Delete(id string) error {
	/* Primero debemos generar una estructura User para poder pasarle el ID a GORM */
	user := User{ID: id}

	result := repo.db.Delete(&user) // El metodo que se usa es el .DELETE

	// Handleamos el error
	if result.Error != nil {
		return result.Error
	}

	// Devolvemos nil. No se devuelve el result
	return nil

}

// Creamos el Metodo UPDATE

func (repo *repo) Update(id string, firstName *string, lastName *string, email *string, phone *string) error {
	values := make(map[string]interface{})

	if firstName != nil {
		values["first_name"] = *firstName
	}

	if lastName != nil {
		values["last_name"] = *lastName
	}

	if email != nil {
		values["email"] = *email
	}

	if phone != nil {
		values["phone"] = *phone
	}

	if result := repo.db.Model(&User{}).Where("id = ?", id).Updates(values); result.Error != nil {
		return result.Error
	}

	return nil
}
