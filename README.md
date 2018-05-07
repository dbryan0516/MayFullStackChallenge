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
curl -X POST http://10.0.2.15:12345/setData/5           //must be integer
curl -X POST http://10.0.2.15:12345/getTransaction/0xa92613e8438deb09913c1842dd230c4038171a706944c0cf87d6d89c0d59a521
```
