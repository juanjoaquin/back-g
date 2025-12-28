package enrollment

import (
	"log"

	"github.com/juanjoaquin/back-g/internal/domain"
	"gorm.io/gorm"
)

type (
	Repository interface {
		Create(enroll *domain.Enrollment) error
	}

	repo struct {
		db  *gorm.DB
		log *log.Logger
	}
)

func NewRepo(l *log.Logger, db *gorm.DB) Repository {
	return &repo{
		log: l,
		db:  db,
	}
}

func (repo *repo) Create(enroll *domain.Enrollment) error {
	if err := repo.db.Create(enroll).Error; err != nil {
		repo.log.Printf("error: %v", err)
		return err
	}
	repo.log.Println("enrollment created with id:", enroll.ID)
	return nil
}
