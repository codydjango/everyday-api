package main

import (
	"net/http"
)

// Route routes to use
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes are the registered routes
type Routes []Route

var routes = Routes{
	Route{
		"Version",
		"GET",
		"/version/",
		VersionHandler,
	},
	Route{
		"Nonce",
		"GET",
		"/{address}/nonce/",
		AddressNonceHandler,
	},
}
