// https://www.alexedwards.net/blog/golang-response-snippets

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Profile struct {
	Nonce     string `json:"nonce"`
	publicKey string `json:"publicKey"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"Cody8", ""}

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

	// fmt.Fprintf(w, "Hello World from path: %s\n", r.URL.Path)
}

func main() {
	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "3001"
	}

	http.HandleFunc("/", Index)
	fmt.Printf("Running on localhost:%s", PORT)
	http.ListenAndServe(":"+PORT, nil)
}
