package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}
var movies []Movie 

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
func main() {
	r := mux.NewRouter()
	
	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie one",Director: &Director{Firstname: "Omar", Lastname: "Ghandour"}})
	movies = append(movies, Movie{ID: "2", Isbn: "438228", Title: "Movie two",Director: &Director{Firstname: "Hamza", Lastname: "Kelany"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	// r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	// r.HandleFunc("/movies", createMovies).Methods("POST")
	// r.HandleFunc("/movies/{id}", updateMovies).Methods("PUT")
	// r.HandleFunc("/movies/{id}", deleteMovies).Methods("DELETE")
fmt.Printf("starting server at port 8000\n")
log.Fatal(http.ListenAndServe(":8000", r))
}