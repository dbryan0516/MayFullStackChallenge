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

	"math/big"
	"context"
	"github.com/dbryan0516/MayFullStackChallenge/contract"
)

type Settings struct {
	Keystore	string `json:"keystore"` // keystore file
	Password	string `json:"password"` // password for keystore file
	NodeURL		string `json:"nodeUrl"`  // URL of an ethereum node ex. "https://10.10.10.10:12345"
}

var settings Settings
var simpleStorage contract.SimpleStorage

// This should take parameters or read from a file
// but I felt it simpler to hardcode these values for the time being
func initSettings() {
	keystoreFile := "keystore.key"  	// keystore file for ETH address
	password := "quorumtest"        	// hardcoded for now
	nodeUrl := "http://localhost:22000" // taken from the cli argument when creating constellation nodes

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

// returns bind.TransactOpts is the authentication needed to deploy contract
func getAuthentication() (*bind.TransactOpts, error){
	//check that the keystore file exists
	file, err := os.Open(settings.Keystore)
	if err != nil {
		fmt.Printf(fmt.Sprintf("File: %s\n", settings.Keystore))
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
	//identify(r)
	//TODO: implement function
	w.Write([]byte("Not Implemented Yet\n"))
}

func deployContract(w http.ResponseWriter, r *http.Request){
	//identify(r)

	conn, err := connectToNode()
	if err != nil {
		w.Write([]byte(err.Error() + "\n"))
		return
	}

	auth, err := getAuthentication()
	if err != nil {
		w.Write([]byte(err.Error() + "\n"))
		return
	}

	address, tx, ssContract, err := contract.DeploySimpleStorage(auth, conn)
	if err != nil {
		w.Write([]byte(err.Error() + "\n"))
		return
	} else if *ssContract == (contract.SimpleStorage{}) {
		w.Write([]byte("Error deploying contract: Contract empty"))
		return
	}

	// save the contract and info for further use
	simpleStorage = *ssContract

	//write response with address and transaction
	response := fmt.Sprintf("{address: \"0x%x\", transactionId: \"0x%x\"}", address, tx.Hash())
	w.Write([]byte(response + "\n"))
}

func setData(w http.ResponseWriter, r *http.Request){
	//identify(r)

	vars := mux.Vars(r)
	if vars == nil || len(vars) != 1 {
		w.Write([]byte("Incorrect number of args passed, expecting 1 integer"))
		return
	}

	// check that arg is an integer/big.Int
	n := new(big.Int)
	n, ok := n.SetString(vars["integer"], 10)
	if !ok {
		w.Write([]byte("Incorrect argument passed, expecting integer formatted string"))
		return
	}

	auth, err := getAuthentication()
	if err != nil {
		w.Write([]byte(err.Error() + "\n"))
	}

	//set the value on the contract
	tx, err := simpleStorage.Set(auth, n)
	if err != nil {
		w.Write([]byte("Incorrect argument passed, expecting integer formatted string"))
		return
	}

	//return the transactionId
	response := fmt.Sprintf("{transactionId: \"0x%x\"}", tx.Hash())
	w.Write([]byte(response + "\n"))

}

func getData(w http.ResponseWriter, r *http.Request){
	//identify(r)

	//set the value on the contract
	opts := bind.CallOpts{Pending: true}
	stored, err := simpleStorage.Get(&opts)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	//return the transactionId
	response := fmt.Sprintf("{storedData: \"%s\"}", stored.String())
	w.Write([]byte(response + "\n"))
}

func getTransaction(w http.ResponseWriter, r *http.Request){
	//identify(r)
	vars := mux.Vars(r)
	if vars == nil || len(vars) != 1 {
		w.Write([]byte("Incorrect number of args passed, expecting 1"))
		return
	}

	conn, err := connectToNode()
	if err != nil {
		w.Write([]byte(err.Error() + "\n"))
		return
	}

	//convert from hexadecimal to hash for lookup
	// TODO: not sure how to error check here
	hash := common.HexToHash(vars["id"])

	//get the transaction by the hash
	// TODO: not sure what context to use but Background doesnt seem like what I want
	tx, pending, err := conn.TransactionByHash(context.TODO(), hash)
	if err != nil {
		// No transaction found
		w.Write([]byte(err.Error() + "\n"))
		return
	}

	chainId := tx.ChainId()
	gas := tx.Gas() 		//gas limit not actual gas

	//src: https://github.com/ethereum/go-ethereum/issues/15069
	var signer types.Signer
	v, _,_ := tx.RawSignatureValues()
	if v.Sign() != 0 && tx.Protected() {
		signer = types.NewEIP155Signer(chainId)
	} else {
		signer = types.HomesteadSigner{}
	}

	sender, err := types.Sender(signer, tx)
	if err != nil {
		w.Write([]byte(err.Error() + "\n"))
		return
	}

	var pendingString string
	if pending {
		pendingString = "true"
	} else {
		pendingString = "false"
	}

	response := fmt.Sprintf("{transactionId: \"0x%x\", pending: \"%s\", gasLimit: \"%d\", sender: \"%s\"}", hash, pendingString, gas, sender.String())
	w.Write([]byte(response + "\n"))
}

func main(){
	initSettings()

	r := mux.NewRouter()
	//in a larger system these endpoints would have a different naming convention
	//but for the purpose of this exercise, I felt these were good enough
	r.HandleFunc("/login", login).Methods("POST")
	r.HandleFunc("/deployContract", deployContract).Methods("POST")
	r.HandleFunc("/setData/{integer}", setData).Methods("POST")
	r.HandleFunc("/getData", getData).Methods("GET")
	r.HandleFunc("/getTransaction/{id}", getTransaction).Methods("POST")
	log.Fatal(http.ListenAndServe(":12345", r))
}
