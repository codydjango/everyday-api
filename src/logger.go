package main

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		start := time.Now()

		next.ServeHTTP(responseWriter, request)

		log.Printf(
			"%s\t%s\t%s\t%s",
			request.Method,
			request.RequestURI,
			name,
			time.Since(start),
		)
	})
}
