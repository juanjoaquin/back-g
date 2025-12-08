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

	/*
		5. Vamos a generar un struct para los errores de las response:
	*/
	ErrorRes struct {
		Error string `json:"error"`
	}
)

// 3. Esta es la función de MakeEndpoints, que va a devolver una estructura de Edpoints. Estos son los que vamos a poder utilizar en nuestro dominio.

// Ahora le pasaremos el Service. Este lo tendra como prop. También lo recibira todas las funciones que encapsula.
func MakeEndpoints(s Service) Endpoints {
	// Returnamos los endpoints
	return Endpoints{
		// Debemos indicar que cada endpoint representa cada funcion
		Create: makeCreateEndpoint(s),
		GetAll: makeGetAllEndpoint(s),
		Get:    makeGetEndpoint(s),
		Update: makeUpdateEndpoint(s),
		Delete: makeDeleteEndpoint(s),
	}
}

// Estas seran una funcion privada, ya que empiezan con minuscula, porque el que vamos a usar es el de arriba
func makeDeleteEndpoint(s Service) Controller {
	// Definimos la funcion del Controller, que seria la que esta arriba de todo del Controller
	return func(w http.ResponseWriter, r *http.Request) {
		// Aqui ira nuestra logica del endpoint
		fmt.Println("delete user")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}

// Create Endpoint
// Aqui tambien le pasaremos ese servicio
func makeCreateEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {

		// Aca llamamos el struct que hicimos del CreateReq (que se encuentra dentro del Type)
		var req CreateReq

		// Con esto inyectamos los datos que vienen del JSON a mi struct (Ej: "first_name":"Nahuel" CreateReq { FirstName: "Nahuel"})
		// Tenemos que hacer un matcheo con lo que trajo el request. Esto lo hacemos con el package json para decodificar la Request.
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			// Con esto manejamos el error. Debemos utilizar el parametro: w http.ResponseWriter, que tenemos ahi en el param
			w.WriteHeader(400)
			// 6. Aca le pasamos el Error struct que hicimos previamente
			json.NewEncoder(w).Encode(ErrorRes{"Invalid request format"})
			return
		}

		// Para pasarlo como Bad Request a uno de los campos, debemos hacerlo asi:
		if req.FirstName == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(ErrorRes{"First name is required"})
			return
		}

		// Vamos a returnar la capa de Servicio que tenemos. En este caso sería: s.Create() con el Body que le habiamos pasado.
		err := s.Create(req.FirstName, req.LastName, req.Phone, req.Email)
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(ErrorRes{err.Error()})
		}

		json.NewEncoder(w).Encode(req)
	}
}

// Get All Endpoint
func makeGetAllEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("get all user")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}

// Get by id endpoint
func makeGetEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("get user")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}

// Update endpoint
func makeUpdateEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("update user")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}
