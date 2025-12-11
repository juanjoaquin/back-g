package user

import (
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Create(user *User) error // Le pasamos como puntero al User
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
