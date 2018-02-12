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
var routes = []Route{
	Route{
		"Index",
		"GET",
		"/v1/",
		Index,
	},
	Route{
		"Infra",
		"GET",
		"/v1/infra",
		InfraRoutes,
	},
	Route{
		"Infra",
		"GET",
		"/v1/app",
		AppRoutes,
	},
	Route{
		"Infra",
		"GET",
		"/v1/app/docker",
		DockerRoutes,
	},
	Route{
		"Providers",
		"GET",
		"/v1/infra/providers",
		ProvidersIndex,
	},
	Route{
		"ProvidersAWS",
		"GET",
		"/v1/infra/providers/aws",
		ProvidersAWS,
	},
	Route{
		"ResourceAWS",
		"GET",
		"/v1/infra/providers/aws/{ResourceName}",
		ResourceAWS,
	},
	Route{
		"ProvidersChef",
		"GET",
		"/v1/providers/chef",
		ProvidersChef,
	},
	Route{
		"ResourceChef",
		"GET",
		"/v1/providers/chef/{ResourceName}",
		ResourceChef,
	},
	Route{
		"ProvidersOpenstack",
		"GET",
		"/v1/infra/providers/openstack",
		ProvidersOpenstack,
	},
	Route{
		"ResourceOpenstack",
		"GET",
		"/v1/infra/providers/openstack/{ResourceName}",
		ResourceOpenstack,
	},
	Route{
		"Json2Hcl",
		"POST",
		"/v1/json2hcl",
		JSON2Hcl,
	},
}
