package course

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/juanjoaquin/back-g/internal/domain"
	"gorm.io/gorm"
)

type Repository interface {
	Create(course *domain.Course) error
	Get(id string) (*domain.Course, error)
	GetAll(filters Filters, offset int, limit int) ([]domain.Course, error)
	Delete(id string) error
	Count(filters Filters) (int, error)
	Update(id string, name *string, startDate, endDate *time.Time) error
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

// POST: Crear Course
func (repo *repo) Create(course *domain.Course) error {
	if err := repo.db.Create(course).Error; err != nil {
		repo.log.Printf("error: %v", err)
		return err
	}
	repo.log.Println("Course created with id:", course.ID)
	return nil
}

// GET: Get All Courses
func (repo *repo) GetAll(filters Filters, offset, limit int) ([]domain.Course, error) {
	var c []domain.Course

	log.Println("REPO => ejecutando query...")

	tx := repo.db.Model(&domain.Course{})
	tx = applyFilters(tx, filters)

	tx = tx.Limit(limit).Offset(offset)

	result := tx.Order("created_at desc").Find(&c)

	if result.Error != nil {
		return nil, result.Error
	}

	return c, nil
}

func (repo *repo) Count(filters Filters) (int, error) {
	var count int64
	tx := repo.db.Model(domain.Course{})
	tx = applyFilters(tx, filters)
	if err := tx.Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func (repo *repo) Get(id string) (*domain.Course, error) {
	course := domain.Course{ID: id}

	result := repo.db.First(&course)

	if result.Error != nil {
		return nil, result.Error
	}

	return &course, nil
}

func (repo *repo) Delete(id string) error {
	course := domain.Course{ID: id}

	result := repo.db.Delete(&course)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *repo) Update(id string, name *string, startDate *time.Time, endDate *time.Time) error {
	values := make(map[string]interface{})

	if name != nil {
		values["name"] = *name
	}

	if startDate != nil {
		values["start_date"] = *startDate
	}

	if endDate != nil {
		values["end_date"] = *endDate
	}

	if result := repo.db.Model(&domain.Course{}).Where("id = ?", id).Updates(values); result.Error != nil {
		return result.Error
	}

	return nil

}

// FUNCION PARA EL APLICADO DE FILTROS
func applyFilters(tx *gorm.DB, filters Filters) *gorm.DB {

	if filters.Name != "" { // Basicamente que si viene vacio, no pasa nada, y que lo devuelva en lower o uppercase
		filters.Name = fmt.Sprintf("%%%s%%", strings.ToLower(filters.Name))
		tx = tx.Where("lower(name) like ?", filters.Name) // Query de GORM para la consulta
	}

	return tx
}
