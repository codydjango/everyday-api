/*
Create the router, assign the registered API routes, and add a few non-API routes.
*/

package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// CreateRouter creates a router for the api.
func CreateRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}).Methods("GET")

	apiRouter := router.
		PathPrefix("/api").
		Subrouter()

	for _, route := range routes {
		apiRouter.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(Logger(route.HandlerFunc, route.Name))
	}

	return router
}
