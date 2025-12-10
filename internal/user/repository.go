package user

import (
	"log"

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
	return nil
}
