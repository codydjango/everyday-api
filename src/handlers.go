/*
Controllers.
*/

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

// VersionHandler controller returns version of the API
func VersionHandler(responseWriter http.ResponseWriter, request *http.Request) {
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

// AddressNonceHandler controller
func AddressNonceHandler(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	publicKey := vars["address"]

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

// AuthenticationHandler controller
func AuthenticationHandler(responseWriter http.ResponseWriter, request *http.Request) {
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

// SessionGetHandler controller
func SessionGetHandler(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	publicKey := vars["address"]
	name := AddressNameLookup[publicKey]
	session := Session{name}

	json, err := json.Marshal(session)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json; charset=UTF-8")
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write(json)
}

// SessionPostHandler controller
func SessionPostHandler(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	publicKey := vars["address"]

	session, err := DecodeSession(request.Body)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	AddressNameLookup[publicKey] = session.Name

	responseWriter.Header().Set("Content-Type", "application/json; charset=UTF-8")
	responseWriter.WriteHeader(http.StatusOK)
}
