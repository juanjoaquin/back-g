package course

import (
	"log"

	"gorm.io/gorm"
)

type Repository interface {
	Create(course *Course) error
}

type repo struct {
	log *log.Logger
	db  *gorm.DB
}

func NewRepo(log *log.Logger, db *gorm.DB) Repository {
	return &repo{
		log: log,
		db:  db,
	}
}

func (repo *repo) Create(course *Course) error {
	if err := repo.db.Create(course).Error; err != nil {
		repo.log.Printf("error: %v", err)
		return err
	}
	repo.log.Println("Course created with id:", course.ID)
	return nil
}
