package main

import (
	"net/http"

	"encoding/json"

	"fmt"

	"github.com/go-chi/chi/v5"
)

type Person struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
}

func main() {

	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("Go web"))
	})

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("Pong"))
	})

	router.Post("/greetings", func(w http.ResponseWriter, r *http.Request) {

		var requestData Person
		err := json.NewDecoder(r.Body).Decode(&requestData)
		if err != nil {
			http.Error(w, "Error al decodificar el JSON", http.StatusBadRequest)
			return
		}

		FirstNameTemp := requestData.FirstName
		LastNameTempo := requestData.LastName

		responseMessage := fmt.Sprintf("Hello %s %s", FirstNameTemp, LastNameTempo)

		w.Write([]byte(responseMessage))

	})

	http.ListenAndServe(":8080", router)

}
