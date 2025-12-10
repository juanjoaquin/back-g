package user

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	Create(user *User) error // Le pasamos como puntero al User
}

// Esta struct va hacer referencia a la DB de GORM
type repo struct {
	db *gorm.DB
}

// Creamos una funcion que va a instanciar este repo.

func NewRepo(db *gorm.DB) Repository {
	return &repo{
		db: db,
	}
}

// Creamos el Metodo Create
func (repo *repo) Create(user *User) error {
	fmt.Println(user)
	return nil
}
