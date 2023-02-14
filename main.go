 package main

 import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	// "gorm.io/gorm"
 )

 type Movie struct{

	ID string `json:"id"`
	IMDB string `json:"imdb"`
	Title string `json:"title"`
	Director *Director `json:"director"`
 }

 type Director struct{
	FirstName string `json:"fname"`
	LastName string `json:"lname"`
 }

 var movies []Movie

 func getMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
 }

func deleteMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index,item := range movies{
		if item.ID == params["id"]{
			fmt.Println("## this is index",index)
			fmt.Println("##### this is item.id",item.ID)
			movies = append(movies[:index], movies[index+1:]... )
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
		
	}
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000))     /// not an good idea to take id random because of redundancy 
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}


func updateMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index,item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index],movies[index+1:]... )
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}

	}
}

 func main(){
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1",IMDB: "5.6",Title: "Titan",Director: &Director{FirstName: "John",LastName: "Cena"}})
	movies = append(movies, Movie{ID: "2",IMDB: "9.8",Title: "Godfather",Director: &Director{FirstName: "Prabhu",LastName: "kumar"}})
	movies = append(movies, Movie{ID: "3",IMDB: "8.8",Title: "IP Man",Director: &Director{FirstName: "xyz",LastName: "abc"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movie/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
 }