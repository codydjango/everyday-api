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
			"http://127.0.0.1:3001",
			"https://codydjango.github.io",
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

	cors := getCors()
	router := cors.Handler(CreateRouter())
	address := fmt.Sprintf("0.0.0.0:%s", PORT)

	srv := &http.Server{
		Addr:         address,
		Handler:      router,
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Second * 10,
	}

	log.Println(fmt.Sprintf("Running on %s", address))
	log.Fatal(srv.ListenAndServe())
}
