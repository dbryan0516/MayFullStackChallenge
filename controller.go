package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
)


func login(w http.ResponseWriter, r *http.Request){
	//TODO: implement function
}

func deployContract(w http.ResponseWriter, r *http.Request){
	//TODO: implement function
}

func setData(w http.ResponseWriter, r *http.Request){
	//TODO: implement function
}

func getData(w http.ResponseWriter, r *http.Request){
	//TODO: implement function
}

func getTransaction(w http.ResponseWriter, r *http.Request){
	//TODO: implement function
}

func main(){
	r := mux.NewRouter()
	//in a larger system these endpoints would have a different naming convention
	//but for the purpose of this exercise, I felt these were good enough
	r.HandleFunc("/login", login).Methods("POST")
	r.HandleFunc("/deployContract", deployContract).Methods("POST")
	r.HandleFunc("/setData", setData).Methods("POST")
	r.HandleFunc("/getData", getData).Methods("GET")
	r.HandleFunc("/getTransaction/{id}", getTransaction).Methods("GET")
	log.Fatal(http.ListenAndServe(":12345", r))
}