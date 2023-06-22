package transport

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go_mongo_test/internal/database"
	"go_mongo_test/internal/domain"
	"go_mongo_test/internal/logic"
	"log"
	"net/http"
	"time"
)

func NewServer() {
	r := mux.NewRouter()
	r.HandleFunc("/cars/add", addCar)
	r.HandleFunc("/cars/getAll", getCars)
	srv := &http.Server{
		Addr:         "127.0.0.1:8080",
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	fmt.Println("Server started")
	log.Fatal(srv.ListenAndServe())
}

func getCars(writer http.ResponseWriter, request *http.Request) {
	cars, err := logic.GetCars()
	if err != nil {
		body, _ := json.Marshal(err)
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(body)
		return
	}
	body, _ := json.Marshal(cars)
	writer.WriteHeader(http.StatusOK)
	writer.Write(body)
	return
}

func addCar(writer http.ResponseWriter, request *http.Request) {
	var c domain.Car
	err := json.NewDecoder(request.Body).Decode(&c)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	var car domain.CarInternal = domain.CarInternal{
		Model:  c.Model,
		Engine: c.Engine,
		Year:   c.Year,
	}

	err = database.AddCar(car)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
