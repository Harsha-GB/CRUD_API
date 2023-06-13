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

type Movie struct{
	Id string `json:"id"`
	Title string `json:"title"`
	Uni_Id string `json:"uni_id"`
	Director *Director `json:"director"` //pointer helps us to access the movie details for that director details
}

type Director struct{
	FirstName string `json:"firstname"`
	LastName string `json:"secondname"`
}

var movies []Movie

func getmovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)

}

func deletemovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params :=mux.Vars(r)

	for i,v :=range movies{
		if v.Id == params["id"]{
			movies=append(movies[:i],movies[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}
func createmovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var movie Movie
	_ =json.NewDecoder(r.Body).Decode(&movie)
	movie.Id=strconv.Itoa(rand.Intn(100000))
	movies=append(movies,movie)
	json.NewEncoder(w).Encode(movie)
}
func updatemovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json") 
	params:=mux.Vars(r)

	for i,v:= range movies{
		if v.Id==params["id"]{
			movies=append(movies[:i],movies[i+1:]...)
			var movie Movie
			_=json.NewDecoder(r.Body).Decode(&movie)
			movie.Id=params["id"]
			movies=append(movies,movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
		
}

func getmovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)

	for _,v :=range movies{
		if v.Id==params["id"]{
			json.NewEncoder(w).Encode(v)
			return
		}

	}
}
func main(){

	r:=mux.NewRouter()
	//First we have add some movies so that initially when run get movies command in postman,we could see some data.

	movies=append(movies,Movie{Id:"1",Title :"Avengers",Uni_Id:"12345",Director :&Director{FirstName:"Stan",LastName:"Lee"}})
	movies=append(movies,Movie{Id:"2",Title :"Avengers_2",Uni_Id:"678910",Director :&Director{FirstName:"Tony",LastName:"stark"}})
	r.HandleFunc("/movies",getmovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getmovie).Methods("GET")
	r.HandleFunc("/movies",createmovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updatemovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deletemovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8080 \n")
	log.Fatal(http.ListenAndServe(":8080",r))
}