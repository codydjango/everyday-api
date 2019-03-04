/*
Controllers.
*/

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// HandleVersion controller returns version of the API
func HandleVersion(responseWriter http.ResponseWriter, request *http.Request) {
	version := "0.0.1"

	keyMapping := make(map[string]string)
	keyMapping["version"] = version

	json, err := json.Marshal(keyMapping)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json; charset=UTF-8")
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write(json)
}

// HandleAuthTest for testing authentication middleware
func HandleAuthTest(responseWriter http.ResponseWriter, request *http.Request) {
	keyMapping := make(map[string]string)
	keyMapping["good"] = "good"

	json, err := json.Marshal(keyMapping)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json; charset=UTF-8")
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write(json)
}

// HandleAddressNonce controller
func HandleAddressNonce(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	publicKey := vars["account"]

	if AddressNonceLookup[publicKey] == 0 {
		AddressNonceLookup[publicKey]++
	}

	keyMapping := make(map[string]string)
	keyMapping["nonce"] = fmt.Sprintf("%d", AddressNonceLookup[publicKey])
	keyMapping["publicKey"] = publicKey

	json, err := json.Marshal(keyMapping)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json; charset=UTF-8")
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write(json)
}

// DecodeClaim is the function that decodes the json into a claim struct
func DecodeClaim(r io.ReadCloser) (x *Claim, err error) {
	x = &Claim{}
	err = json.NewDecoder(r).Decode(x)
	return
}

// HandleAuthentication controller
func HandleAuthentication(responseWriter http.ResponseWriter, request *http.Request) {
	claim, err := DecodeClaim(request.Body)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	if !claim.isValid() {
		responseWriter.Header().Set("Content-Type", "application/json; charset=UTF-8")
		responseWriter.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	claim.updateChallenge()
	claim.verify()

	if claim.Verified == true {
		claim.updateNonce()
		claim.updateToken()
	}

	json, err := json.Marshal(claim)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json; charset=UTF-8")
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write(json)
}

// Session for session info
type Session struct {
	Name string `json:"name,omitempty"`
}

// AddressNameLookup keep track of the nonce for each address
var AddressNameLookup = make(map[string]string)

// DecodeSession is the function that decodes the json into a session struct
func DecodeSession(r io.ReadCloser) (x *Session, err error) {
	x = &Session{}
	err = json.NewDecoder(r).Decode(x)
	return
}

// HandleSessionGet controller
func HandleSessionGet(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	publicKey := vars["account"]
	body := AddressNameLookup[publicKey]

	responseWriter.Header().Set("Content-Type", "application/json; charset=UTF-8")
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write([]byte(body))
}

// HandleSessionPost controller
func HandleSessionPost(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	publicKey := vars["account"]

	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
	}

	AddressNameLookup[publicKey] = string(body)
	responseWriter.Header().Set("Content-Type", "application/json; charset=UTF-8")
	responseWriter.Write([]byte("{}"))
	responseWriter.WriteHeader(http.StatusOK)
}
