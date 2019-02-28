package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
)

func getCors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:1234",
			"https://codydjango.github.io",
		},
		AllowCredentials: true,
		Debug:            false,
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
		Addr:    fmt.Sprintf("0.0.0.0:%s", PORT),
		Handler: cors.Handler(router),
	}

	// WriteTimeout: time.Second * 2,
	// ReadTimeout:  time.Second * 2,
	// IdleTimeout:  time.Second * 5,
	// https://github.com/gorilla/mux#graceful-shutdown

	log.Println(fmt.Sprintf("Running on 0.0.0.0:%s", PORT))
	log.Fatal(srv.ListenAndServe())

	// go func() {
	// 	if err := srv.ListenAndServe(); err != nil {
	// 		log.Println(err)
	// 	}
	// }()

	// var wait time.Duration
	// flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	// flag.Parse()

	// c := make(chan os.Signal, 1)
	// signal.Notify(c, os.Interrupt)
	// <-c

	// ctx, cancel := context.WithTimeout(context.Background(), wait)
	// defer cancel()
	// srv.Shutdown(ctx)
	// log.Println("Shutting down...")
	// os.Exit(0)
}
