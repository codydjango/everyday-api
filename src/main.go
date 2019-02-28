package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
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
		WriteTimeout: time.Second * 2,
		ReadTimeout:  time.Second * 2,
		IdleTimeout:  time.Second * 5,
		Handler:      cors.Handler(router),
	}

	// https://github.com/gorilla/mux#graceful-shutdown
	log.Println(fmt.Sprintf("Running on localhost:%s", PORT))
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("Shutting down...")
	os.Exit(0)
}
