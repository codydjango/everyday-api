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
		HandleVersion,
	},
	Route{
		"Nonce",
		"GET",
		"/account/{account}/nonce/",
		HandleAddressNonce,
	},
	Route{
		"Authentication",
		"POST",
		"/authentication/",
		HandleAuthentication,
	},
}

var authenticatedAPIRoutes = Routes{
	Route{
		"SessionGet",
		"GET",
		"/account/{account}/data/",
		HandleSessionGet,
	},
	Route{
		"SessionPost",
		"POST",
		"/account/{account}/data/",
		HandleSessionPost,
	},
}
