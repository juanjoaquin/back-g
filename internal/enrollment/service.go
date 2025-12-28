package enrollment

import (
	"log"

	"github.com/juanjoaquin/back-g/internal/domain"
)

type (
	Service interface {
		Create(userID, courseID string) (*domain.Enrollment, error)
	}

	service struct {
		log  *log.Logger
		repo Repository
	}
)

func NewService(log *log.Logger, repo Repository) Service {
	return &service{
		log:  log,
		repo: repo,
	}
}

func (s service) Create(userID, courseID string) (*domain.Enrollment, error) {
	enroll := domain.Enrollment{
		UserID:   userID,
		CourseID: courseID,
		Status:   "P",
	}

	if err := s.repo.Create(&enroll); err != nil {
		return nil, err
	}

	return &enroll, nil
}
