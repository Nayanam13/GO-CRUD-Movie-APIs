package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type movie struct {
	ID       string    `jason:"id`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *director `json:"director"`
}

type director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []movie

func main() {
	r := mux.NewRouter()

	movies = append(movies, movie{ID: "1", Isbn: "4626", Title: "movie1", Director: &director{Firstname: "Jonh", Lastname: "Doe"}})
	movies = append(movies, movie{ID: "2", Isbn: "4634", Title: "movie2", Director: &director{Firstname: "Pinky", Lastname: "Master"}})

	r.HandleFunc("/movies", getmovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Server is starting at port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Error occured", err)
	}

}

func getmovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(movies)
	}

}

func getMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "appication/json")
	params := mux.Vars(r)
	for _, item := range movies {

		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}

	}

}

func createMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Movie movie
	_ = json.NewDecoder(r.Body).Decode(&Movie)
	Movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, Movie)
	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var Movie movie
			_ = json.NewDecoder(r.Body).Decode(&Movie)
			Movie.ID = strconv.Itoa(rand.Intn(100000000))
			movies = append(movies, Movie)
		}
	}
}
