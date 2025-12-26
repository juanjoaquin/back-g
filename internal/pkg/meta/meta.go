/*
El package meta basicamente es la informacion que se envia de total pages en un getAll
*/
package meta

import (
	"os"
	"strconv"
)

type Meta struct {
	TotalCount int `json:"total_count"`
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	PageCount  int `json:"page_count"`
}

// Creamos una funcion que va a devolver el objeto Meta
func New(page, perPage, total int) (*Meta, error) {
	/* Page = Total Pages - perPage = Paginacion actual en la que me encuentro */

	// Si perPage viene en 0 le especificamos un perPage por default desde ENV
	if perPage <= 0 {
		var err error
		perPage, err = strconv.Atoi(os.Getenv("PAGINATOR_LIMIT_DEFAULT"))
		if err != nil {
			return nil, err
		}
	}
	// Se debe especificar la cantidad de paginas que yo tengo en el Get All
	pageCount := 0
	//Si el total es mayour o igual a 0, vamos hacer la logica para calcular el page
	if total >= 0 {
		pageCount = (total + perPage - 1) / perPage
		// Si le mandamos mas paginas de la que existen, nos coloque las paginas que realmente hay
		if page > pageCount {
			page = pageCount
		}
	}

	/* Si hay una page en valor negativo, indicamos que nos devuelva la Page 1  */
	if page < 1 {
		page = 1
	}

	// Returnamos al struct pegandole
	return &Meta{
		// Retornar los valores de la page
		Page:       page,
		PerPage:    perPage,
		PageCount:  pageCount,
		TotalCount: total,
	}, nil
}

/* El offset va a ser a partir de que datos, de que numero de filas te voy a traer la informacion*/
func (p *Meta) Offset() int {
	return (p.Page - 1) * p.PerPage
}

/* Limit es hasta el numero de filas */
func (p *Meta) Limit() int {
	return p.PerPage
}
