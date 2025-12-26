package course

import (
	"fmt"
	"log"
	"strings"

	"gorm.io/gorm"
)

type Repository interface {
	Create(course *Course) error
	GetAll(filters Filters, offset int, limit int) ([]Course, error)
	Count(filters Filters) (int, error)
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
func (repo *repo) Create(course *Course) error {
	if err := repo.db.Create(course).Error; err != nil {
		repo.log.Printf("error: %v", err)
		return err
	}
	repo.log.Println("Course created with id:", course.ID)
	return nil
}

// GET: Get All Courses
func (repo *repo) GetAll(filters Filters, offset, limit int) ([]Course, error) {
	var c []Course

	tx := repo.db.Model(c)
	tx = applyFilters(tx, filters)

	tx = tx.Limit(limit).Offset(offset)

	result := tx.Order("created_at desc").Find(&c)

	if result.Error != nil {
		return nil, result.Error
	}

	return c, nil
}

// FUNCION PARA EL APLICADO DE FILTROS
func applyFilters(tx *gorm.DB, filters Filters) *gorm.DB {

	if filters.Name != "" { // Basicamente que si viene vacio, no pasa nada, y que lo devuelva en lower o uppercase
		filters.Name = fmt.Sprintf("%%%s%%", strings.ToLower(filters.Name))
		tx = tx.Where("lower(name) like ?", filters.Name) // Query de GORM para la consulta
	}

	return tx
}

func (repo *repo) Count(filters Filters) (int, error) {
	var count int64
	tx := repo.db.Model(Course{})
	tx = applyFilters(tx, filters)
	if err := tx.Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}
