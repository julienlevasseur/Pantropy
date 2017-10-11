package main

import (
	"net/http"
)

// Route definition
type Route struct {
	Name		string
	Method		string
	Pattern		string `json:"pattern"`
	HandlerFunc http.HandlerFunc
}

// Routes
var routes []Route = []Route{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Providers",
		"GET",
		"/providers",
		ProvidersIndex,
	},
	Route{
		"ProvidersAWS",
		"GET",
		"/providers/aws",
		ProvidersAWS,
	},
	Route{
		"ResourceAWS",
		"GET",
		"/providers/aws/{ResourceName}",
		ResourceAWS,
	},
	Route{
		"ProvidersChef",
		"GET",
		"/providers/chef",
		ProvidersChef,
	},
	Route{
		"ResourceChef",
		"GET",
		"/providers/chef/{ResourceName}",
		ResourceChef,
	},
	Route{
		"ProvidersOpenstack",
		"GET",
		"/providers/openstack",
		ProvidersOpenstack,
	},
	Route{
		"ResourceOpenstack",
		"GET",
		"/providers/openstack/{ResourceName}",
		ResourceOpenstack,
	},
}
