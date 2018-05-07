# MayFullStackChallenge

To be used with the 7-nodes quorum example from [github.com/jpmorgan/quorum-examples](https://github.com/jpmorganchase/quorum-examples)

## Setup
Setup the 7-nodes network as the [README](https://github.com/jpmorganchase/quorum-examples/blob/master/README.md) states.
In the Vagrant VM, run the following commands
Clone this repo with go
```
  go get github.com/dbryan0516/MayFullStackChallenge 
```
Then build the project
```
  cd $GOPATH/src/github.com/dbryan0516/MayFullStackChallenge   //goes into dir
  go build
```
This pulls the necessary repos to build project.
Now you must add a valid Ethereum keystore file to the MayFullStackChallenge directory, named Keystore.key

**This implementation assumes a keystore file with the passphrase "quorumtest"**


Now running: ```go run controller.go &  ``` Runs the api handler in background so the user can execute curl commands in same window
  
  

## Usage Examples
```
curl -X POST http://10.0.2.15:12345/deployContract
  ex. {address: "0xd6a8c1446c7b60f50167b6c3e629e4093c19f275", transactionId: "0xde271415b2db0f34520c0f9adcc477ef905d7a8ece950ca1f8f4fffa8d7f9123"}
curl -X POST http://10.0.2.15:12345/setData/5           //must be integer
  ex. {transactionId: "0x8af368b9be0f6a29c5f74d2e5487c7a1d487f1bac9f24ca3cf760d9826c909fa"}
curl -X POST http://10.0.2.15:12345/getTransaction/0x8af368b9be0f6a29c5f74d2e5487c7a1d487f1bac9f24ca3cf760d9826c909fa
  ex. {transactionId: "0x8af368b9be0f6a29c5f74d2e5487c7a1d487f1bac9f24ca3cf760d9826c909fa", pending: "true", gasLimit: "3754426369", sender: "0x3E7dF3682758318ace84F9E7f1e07b63173d91ff"}
curl -X GET http://10.0.2.15:12345/getData
  ex. {storedData: "5"}
```
