package main

import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct {	
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`

}
var movies []Movie

func main()  {
	r  := mux.NewRouter() // function from gorila

	movies = append(movies,Movie{
		ID: "1", 
		Isbn: "438712", 
		Title: "ONE", 
		Director: &Director{FirstName: "one", LastName: "Three"}}
	movies = append(movies,Movie{
		ID: "2", 
		Isbn: "433342", 
		Title: "Stree", 
		Director: &Director{FirstName: "John", LastName: "Doe"}}
	

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/id", getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/id",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/id",deleteMovie).Methods("DELETE")

	fmt.Println("Starting servr at port /8000")
	log.Fatal(http.ListenAndServe(":8000", nil))

}