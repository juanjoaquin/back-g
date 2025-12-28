package enrollment

import (
	"encoding/json"
	"net/http"

	"github.com/juanjoaquin/back-g/internal/pkg/meta"
)

type (
	Controller func(w http.ResponseWriter, r *http.Request)

	Endpoints struct {
		Create Controller
	}

	CreateReq struct {
		UserID   string `json:"user_id"`
		CourseID string `json:"course_id"`
	}

	Response struct {
		Status int         `json:"status"`
		Data   interface{} `json:"data,omitempty"`
		Err    string      `json:"error,omitempty"`
		Meta   *meta.Meta  `json:"meta,omitempty"`
	}
)

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		Create: makeCreateEndpoint(s),
	}
}

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
			json.NewEncoder(w).Encode(&Response{Status: 400, Err: "Invalid request format"})
			return
		}

		// Para pasarlo como Bad Request a uno de los campos, debemos hacerlo asi:
		if req.UserID == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{Status: 400, Err: "UserID is required"})
			return
		}

		if req.CourseID == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{Status: 400, Err: "CourseID is required"})
			return
		}

		// Vamos a returnar la capa de Servicio que tenemos. En este caso sería: s.Create() con el Body que le habiamos pasado.
		user, err := s.Create(req.UserID, req.CourseID)
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{Status: 400, Err: err.Error()})
		}

		json.NewEncoder(w).Encode(&Response{Status: 201, Data: user})
	}
}
