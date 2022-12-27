package router

import (
	"github.com/WellDoneIT90/mongoapi/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controllers.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controllers.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movies/{id}", controllers.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/deletemovies", controllers.DeleteAllMovies).Methods("DELETE")

	return router
}
