package user

import "log"

// Nuestro servicio lo vamos a menajar con interfaces. Esto nos facilitara para mockearlo, o utilizarlo de forma mas generica
type Service interface {
	/* 	1. Vamos a definirle los metodos de los Endpoints que fuimos utilizando.
	   	Le pasaremos tambien los elementos del body del Create por ejemplo */
	Create(firstName, lastName, email, phone string) error
}

/* 2. Vamos a definir una struct, está sera en privado */
type service struct {
}

/*
 3. Haremos una funcion llamada: NewService
    Esta lo que hara sera crear un nuevo servicio, que esta ser la interface.
*/
func NewService() Service {
	return &service{}
}

/* 4. Vamos a generar un metodo, que esto se lo deberemos pasar a la funcion de NewService. */
func (s service) Create(firstName, lastName, email, phone string) error {
	log.Println("Create user service")
	return nil
}
