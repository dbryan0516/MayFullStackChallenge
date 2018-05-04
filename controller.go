package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"io/ioutil"
	"errors"
	"os"
	"fmt"
)

type Settings struct {
	Keystore string `json:keystore` // contents of keystore file
	Password string `json:password` // password for keystore file
}

func getSettings() (Settings, error){
	keystoreFile := "./keystore.key"  // keystore file for ETH address
	password := "quorumtest"          // hardcoded for now
	var settings Settings

	keystore, err := ioutil.ReadFile(keystoreFile)
	if err != nil {
		return settings, errors.New("error reading keystore file")
	}

	settings.Keystore = string(keystore)
	settings.Password = password

	return settings, nil
}

// gets the identification of the user logged in
func identify(r *http.Request){
	//TODO: implement function
}

func login(w http.ResponseWriter, r *http.Request){
	identify(r)
	//TODO: implement function
}

func deployContract(w http.ResponseWriter, r *http.Request){
	identify(r)
	//TODO: implement function
}

func setData(w http.ResponseWriter, r *http.Request){
	identify(r)
	//TODO: implement function
}

func getData(w http.ResponseWriter, r *http.Request){
	identify(r)
	//TODO: implement function
}

func getTransaction(w http.ResponseWriter, r *http.Request){
	identify(r)
	//TODO: implement function
}

func main(){
	settings, err := getSettings()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	} else if settings == (Settings{}) {
		fmt.Print("Settings incorrect")
		os.Exit(1)
	}

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