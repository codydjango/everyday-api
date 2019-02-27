package main

import (
	"encoding/json"
	"net/http"
	"strconv"

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

// AddressNonceLookup keep track of the nonce for each address
var AddressNonceLookup = make(map[string]int)

// AddressNonceHandler controller
func AddressNonceHandler(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	publicKey := vars["address"]
	AddressNonceLookup[publicKey]++
	keyMapping := make(map[string]string)
	keyMapping["nonce"] = strconv.Itoa(AddressNonceLookup[publicKey])
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