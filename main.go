package main

/*
La arquitectura sera por capas.

- Presentation Layer (Capa de presentacion)
Esta va a ser el endpoint, el controller. Va a recibir el request y va a validar que los campos que son requeridos
que no vengan en blanco. Estas validaciones seran pegando al request.

El Endpoint le va a pasar la responsabilidad a la capa de servicio, que va a tener toda la logica de negocio
- Business Layer (Capa de negocio)
Va a tener toda la logica de negocio. El servicio es el que mas cargado estara, que por ahi viene una request de
un determinado body, y lo que puede llegar hacer, es que los datos, que viene de un campo con otro servicio, lo valide
con ese otro servicio.

Una vez que el servicio haga esto, le pasara la responsabilidad a la capa de Persistencia.
- Persistance Layer (Capa de persistencia)
Esta maneja la informacion con la base de datos, es un comunicador desde la base de datos. Hace el create, el update,
genera un id.

- Database Layer (Capa de base de datos)
Es la base de datos
*/

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/juanjoaquin/back-g/internal/user" // Importamos el package user
)

func main() {
	// 2. Aca importamos el package http, y la funcion HandleFunc. Con esto definimos el path que vamos a utilizar para getear a la funcion.
	// http.HandleFunc("/users", getUsers)
	// http.HandleFunc("/courses", getCourses)

	router := mux.NewRouter()

	userEndpoint := user.MakeEndpoints()

	router.HandleFunc("/users", userEndpoint.GetAll).Methods("GET")
	router.HandleFunc("/users", userEndpoint.Create).Methods("POST")
	router.HandleFunc("/users", userEndpoint.Update).Methods("PATCH")

	//3. Vamos a decirle, que levante y escuche y levante el PUERTO quer le vamos a pasar.
	// Con esto usamos la funcion: http.ListenAndServe(port). Y aqui entre las (port) le pasamos el puerto.
	/* 	err := http.ListenAndServe(port, nil)
	   	if err != nil {
	   		fmt.Println("Error:", err)
	   	} */

	// 4. Vamos a usar el package de Gorilla Mux, que es basicamente un router handler de Go mucho mejor
	// Para eso vamos hacer un: go get https://github.com/gorilla/mux

	/*
		5. Vamos a crear una carpeta llamada "internal".
		En esta tendremos todos los packages internos propios. No lo trabajaremos con packages externos. Para ello debemos crear una folder llamada "pkg".
		Dentro de internal, respetamos todos los procesos internos, ya sea interno de nuestro proyecto, o externos que nos pertenezcan.
	*/

	// 6. Creamos nuestro servidor para poder levantarlo.
	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println("Error:", err)
	}

}

// func getUsers(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("got /users")
// 	/* 1. Vamos a sobreescribir, modificar, el response, para que me devuelva una respuesta que yo le quiera definir */
// 	/* Para eso usamos el package IO. y usamos la funcion WriteString. En este caso, la reescribimos a la response con el string que pasamos */
// 	io.WriteString(w, "This is my user endpoint \n")
// }

/* func getCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got /courses")
	io.WriteString(w, "el lolo \n")

} */
