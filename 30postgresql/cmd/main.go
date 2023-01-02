package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/WellDoneIT90/Golang_project/30postgresql/pkg/db"
	"github.com/WellDoneIT90/Golang_project/30postgresql/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("PostgreSQL in Golang")

	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/books", h.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", h.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books", h.AddBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}", h.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books/{id}", h.DeleteBook).Methods(http.MethodDelete)

	log.Println("API is runnung!")
	http.ListenAndServe(":4000", router)
}
