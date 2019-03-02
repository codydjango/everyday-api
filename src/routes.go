/*
API Routes
*/

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

var apiRoutes = Routes{
	Route{
		"Version",
		"GET",
		"/version/",
		VersionHandler,
	},
	Route{
		"Nonce",
		"GET",
		"/address/{address}/nonce/",
		AddressNonceHandler,
	},
	Route{
		"Authentication",
		"POST",
		"/authentication/",
		AuthenticationHandler,
	},
}

var authenticatedAPIRoutes = Routes{
	Route{
		"SessionGet",
		"GET",
		"/address/{address}/session/",
		SessionGetHandler,
	},
	Route{
		"SessionPost",
		"POST",
		"/address/{address}/session/",
		SessionPostHandler,
	},
}
