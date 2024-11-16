package router

import (
	"github.com/gorilla/mux"
	"github.com/tanmay958/dbapi/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()
    router.HandleFunc("/api/movies",controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movie",controller.CreateMovie).Methods("POST") 
	router.HandleFunc("/api/movie/{id}",controller.MarskAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}",controller.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteall",controller.DeleteAllMovie).Methods("DELETE")
	 


	return router 
}