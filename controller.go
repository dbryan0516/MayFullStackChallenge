package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"os"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"./contract"
)

type Settings struct {
	Keystore	string `json:"keystore"` // keystore file
	Password	string `json:"password"` // password for keystore file
	NodeURL		string `json:"nodeUrl"`  // URL of an ethereum node ex. "https://10.10.10.10:12345"
}

var settings Settings
var simpleStorage contract.SimpleStorage
var contractAddress common.Address
var deploymentTx types.Transaction


// This should take parameters but I felt it simpler to hardcode these values for the time being
func initSettings() {
	var settings Settings
	keystoreFile := "./keystore.key"  	// keystore file for ETH address
	password := "quorumtest"        	// hardcoded for now
	nodeUrl := "https://127.0.0.1:9001" // taken from the cli argument when creating constellation nodes

	settings.Keystore = keystoreFile
	settings.Password = password
	settings.NodeURL = nodeUrl
}


// returns the Client pointer *ethclient.Client implements the bind.contractBackend interface, needed to deploy contract
func connectToNode() (*ethclient.Client, error){
	client, err := ethclient.Dial(settings.NodeURL)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// bind.TransactOpts is the authentication needed to deploy contract
func getAuthentication() (*bind.TransactOpts, error){
	//check that the keystore file exists
	file, err := os.Open(settings.Keystore)
	if err != nil {
		log.Printf("Failed to open keystore file: %v\n", err)
		return nil, err
	}

	// create transaction signer from keystore and password
	trans, err := bind.NewTransactor(file, settings.Password)
	if err != nil {
		return nil, err
	}

	return trans, nil
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

	conn, err := connectToNode()
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	auth, err := getAuthentication()
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	address, tx, contract, err := contract.DeploySimpleStorage(auth, conn)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	// save the contract and info for further use
	simpleStorage = *contract
	contractAddress = address
	deploymentTx = *tx

	//write response with address and transaction
	response := fmt.Sprintf("{address: \"0x%x\", transaction: \"0x%x\"}", address, tx.Hash())

	w.Write([]byte(response))
}

func setData(w http.ResponseWriter, r *http.Request){
	identify(r)
	vars := mux.Vars(r)
	if vars == nil {
		//error no parameters in POST call
	}

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
	initSettings()

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