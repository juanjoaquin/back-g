package enrollment

import (
	"errors"
	"log"

	"github.com/juanjoaquin/back-g/internal/course"
	"github.com/juanjoaquin/back-g/internal/domain"
	"github.com/juanjoaquin/back-g/internal/user"
)

type (
	Service interface {
		Create(userID, courseID string) (*domain.Enrollment, error)
	}

	/* Para la validacion del UserID y del CourseID en la creación. Nos debemos traer sus respectivos Services (user.service.go) */
	service struct {
		log        *log.Logger
		repo       Repository
		userSrv    user.Service   // User Service para la validacion
		courseServ course.Service // Course Service para la validacion
	}
)

// Hay que instanciar el Course & User Service en el NewService
func NewService(log *log.Logger, userSrv user.Service, courseSrv course.Service, repo Repository) Service {
	return &service{
		log:        log,
		userSrv:    userSrv,
		courseServ: courseSrv,
		repo:       repo,
	}
}

func (s service) Create(userID, courseID string) (*domain.Enrollment, error) {
	enroll := domain.Enrollment{
		UserID:   userID,
		CourseID: courseID,
		Status:   "P",
	}

	if _, err := s.userSrv.Get(enroll.UserID); err != nil {
		return nil, errors.New("User ID doesnt exists")
	}

	if _, err := s.courseServ.Get(enroll.CourseID); err != nil {
		return nil, errors.New("Course ID doesnt exists")
	}

	if err := s.repo.Create(&enroll); err != nil {
		return nil, err
	}

	return &enroll, nil
}
