// Aqui vamos a generar nuestros endpoints

// 1. Vamos a crear una funcion llamada "MakeEndpoints". Esta se encargara de crear nuestros endpoints
// 2. Vamos a crear una struct, que va a tener todos los endpoints que nosotros vayamos a utilizar
package user

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	// Pasamos al Controller que este recibira un ResponseWritter, y un Request, todo esto lo recibiran los endpoints
	Controller func(w http.ResponseWriter, r *http.Request)

	Endpoints struct {
		// Aqui definimos los endpoints:
		Create Controller
		Get    Controller
		GetAll Controller
		Update Controller
		Delete Controller
	}

	/* 	4. Vamos a definir nuestro request para arrancar.
	   	Vamos a crear una Struct donde vamos a tener los campos que vamos a recibir.
		Esto se lo debemos pasar al controlador que tenemos en el Create, de abajo.

	*/
	CreateReq struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
	}
)

// 3. Esta es la función de MakeEndpoints, que va a devolver una estructura de Edpoints. Estos son los que vamos a poder utilizar en nuestro dominio.
func MakeEndpoints() Endpoints {
	// Returnamos los endpoints
	return Endpoints{
		// Debemos indicar que cada endpoint representa cada funcion
		Create: makeCreateEndpoint(),
		GetAll: makeGetAllEndpoint(),
		Get:    makeGetEndpoint(),
		Update: makeUpdateEndpoint(),
		Delete: makeDeleteEndpoint(),
	}
}

// Estas seran una funcion privada, ya que empiezan con minuscula, porque el que vamos a usar es el de arriba
func makeDeleteEndpoint() Controller {
	// Definimos la funcion del Controller, que seria la que esta arriba de todo del Controller
	return func(w http.ResponseWriter, r *http.Request) {
		// Aqui ira nuestra logica del endpoint
		fmt.Println("delete user")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}

// Create Endpoint
func makeCreateEndpoint() Controller {
	return func(w http.ResponseWriter, r *http.Request) {

		// Aca llamamos el struct que hicimos del CreateReq (que se encuentra dentro del Type)
		var req CreateReq

		// Con esto inyectamos los datos que vienen del JSON a mi struct (Ej: "first_name":"Nahuel" CreateReq { FirstName: "Nahuel"})
		// Tenemos que hacer un matcheo con lo que trajo el request. Esto lo hacemos con el package json para decodificar la Request.
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			// Con esto manejamos el error. Debemos utilizar el parametro: w http.ResponseWriter, que tenemos ahi en el param
			w.WriteHeader(400)
			return
		}

		json.NewEncoder(w).Encode(req)
	}
}

// Get All Endpoint
func makeGetAllEndpoint() Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("get all user")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}

// Get by id endpoint
func makeGetEndpoint() Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("get user")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}

// Update endpoint
func makeUpdateEndpoint() Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("update user")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}
