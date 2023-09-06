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
		Director: &Director{FirstName: "one", LastName: "Three"}})
	movies = append(movies,Movie{
		ID: "2", 
		Isbn: "433342", 
		Title: "Stree", 
		Director: &Director{FirstName: "John", LastName: "Doe"}})
	

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/id", getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/id",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/id",deleteMovie).Methods("DELETE")

	fmt.Println("Starting servr at port /8000")
	log.Fatal(http.ListenAndServe(":8000", nil))

}

func getMovies(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")  //Setting it bcz we want to work with json
	json.NewEncoder(w).Encode(movies)  // this will encode the json data from move slice
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	// mux.Vars is used to retireve route variables i.e.
	// url is localhost:8080/movie/id so to get id we are using this
	params := mux.Vars(r) // here params contains the id 
	// ranging over all the movies
	for index, item := range movies {

		// index is for the slice and item is each movie structure
		// accessing the movie slice using the index
		if item.ID == params["id"] {
            movies = append(movies[:index], movies[index+1:]...) //whatever index the movie is in its index is being updated
            break
        }
    }
	// returning the rmaining moveis
	json.NewEncoder(w).Encode(movies)
}

	
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	
	for _, item := range movies {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}

func createMovie(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	var movie Movie

	// sending the data in the body as it is large and url does not have any parts of it
	_ = json.NewDecoder(r.Body).Decode(&movie)
	// getting new movie id from random function which is integer
	movie.ID = strconv.Itoa(rand.Intn(100000000)) // converting to string
	movies = append(movies,movie)
	json.NewEncoder(w).Encode(movie)

}

func updateMovie(w http.ResponseWriter, r *http.Request)  {
	// delete the movie id which we have sent and append new movie with same id
	w.Header().Set("Content-Type","application/json")
    
    params := mux.Vars(r)
    // range over movies and delete the movie with the id
	// add new movie we sent in the body
	for index, item := range movies {
        if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...) // deleteing the movie
			// appending the movie
			var movie Movie
            _ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
            json.NewEncoder(w).Encode(movie)
            return
        }
    }
}