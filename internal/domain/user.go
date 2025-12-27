package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// 1. Vamos a definir la struct de User con estos campos del Domain, que esta era la capa que interactuaba con la base de datos.
// Pero evitaba que interactue directamente sin capas entre la db y las request. Esta es la tabla intermedia
type User struct {
	ID        string     `json:"id" gorm:"type:char(36);not null;primary_key;unique_index"` // 2. Al usar GORM como ORM, debemos definirle los valores que posee la base de datos. Y sus limites, y demas.
	FirstName string     `json:"first_name" gorm:"type:char(50);not null"`
	LastName  string     `json:"last_name" gorm:"type:char(50);not null"`
	Email     string     `json:"email" gorm:"type:char(50);not null"`
	Phone     string     `json:"phone" gorm:"type:char(30);not null"`
	Course    *Course    `gorm:"-"`
	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
	// Agregamos el campo de Deleted, que es como un Soft Delete de GORM.
	// Nos muestra un registro de los Soft Deleted Users
	Deleted gorm.DeletedAt `json:"-"`
}

// Esta función genera un UUID al crear un User. Anteriormente esto lo hacia el Repository
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	return
}
