// https://www.alexedwards.net/blog/golang-response-snippets

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/rs/cors"
)

func getCors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:1234",
			"http://codydjango.github.io",
			"192.30.252.153",
			"192.30.252.154",
		},
		AllowCredentials: true,
		Debug:            true,
	})
}

func main() {
	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "3001"
	}

	router := NewRouter()
	cors := getCors()

	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%s", PORT),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      cors.Handler(router),
	}

	log.Printf("Running on localhost:%s", PORT)
	log.Fatal(srv.ListenAndServe())
}
