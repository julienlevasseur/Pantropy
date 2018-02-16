package main

import (
	"net/http"
)

// Route definition
type Route struct {
	Name        string
	Method      string
	Pattern     string `json:"pattern"`
	HandlerFunc http.HandlerFunc
}

// Routes
var routes = []Route{
	{
		"Index",
		"GET",
		"/v1/",
		Index,
	},
	{
		"Providers",
		"GET",
		"/v1/providers",
		ProvidersIndex,
	},
	{
		"ProvidersAWS",
		"GET",
		"/v1/providers/aws",
		ProvidersAWS,
	},
	{
		"ResourceAWS",
		"GET",
		"/v1/providers/aws/{ResourceName}",
		ResourceAWS,
	},
	{
		"ProvidersChef",
		"GET",
		"/v1/providers/chef",
		ProvidersChef,
	},
	{
		"ResourceChef",
		"GET",
		"/v1/providers/chef/{ResourceName}",
		ResourceChef,
	},
	{
		"ProvidersOpenstack",
		"GET",
		"/v1/providers/openstack",
		ProvidersOpenstack,
	},
	{
		"ResourceOpenstack",
		"GET",
		"/v1/providers/openstack/{ResourceName}",
		ResourceOpenstack,
	},
	{
		"Json2Hcl",
		"POST",
		"/v1/json2hcl",
		JSON2Hcl,
	},
}
