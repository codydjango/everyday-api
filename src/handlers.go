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
	// x = new(Claim)
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
