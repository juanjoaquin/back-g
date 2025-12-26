/*
El package meta basicamente es la informacion que se envia de total pages en un getAll
*/
package meta

type Meta struct {
	TotalCount int `json:"total_count"`
}

// Creamos una funcion que va a devolver el objeto Meta
func New(total int) (*Meta, error) {

	// Returnamos al struct pegandole
	return &Meta{
		TotalCount: total,
	}, nil
}
