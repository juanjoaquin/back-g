package course

import (
	"log"
	"time"
)

type Service interface {
	Create(name, startDate, endDate string) (*Course, error)
	GetAll(filters Filters, offset, limit int) /* Pasamos el Filtrado de params */ ([]Course, error) // Get All
	Count(filters Filters) (int, error)
}

type Filters struct {
	Name string
}

type service struct {
	log *log.Logger
	// Ahora debemos pasar el Repository
	repo Repository
}

func NewService(log *log.Logger, repo Repository) Service {
	return &service{
		log:  log,
		repo: repo,
	}
}

func (s service) Create(name, startDate, endDate string) (*Course, error) {

	// Parseamos para los valores de Fecha
	startDateParsed, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		s.log.Println(err)
		return nil, err
	}

	endDateParsed, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		s.log.Println(err)
		return nil, err
	}

	course := Course{
		Name:      name,
		StartDate: startDateParsed,
		EndDate:   endDateParsed,
	}

	if err := s.repo.Create(&course); err != nil {
		return nil, err
	}

	return &course, nil
}

func (s service) GetAll(filters Filters, offset, limit int) ([]Course, error) {
	courses, err := s.repo.GetAll(filters, limit, limit)
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func (s service) Count(filters Filters) (int, error) {
	return s.repo.Count(filters)
}
