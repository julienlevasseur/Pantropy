package pantropy

import (
	"net/http"
)

type Route struct {
	Name		string
	Method		string
	Pattern		string `json:"pattern"`
	HandlerFunc http.HandlerFunc
}

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
		"ProvidersChef",
		"GET",
		"/providers/chef",
		ProvidersChef,
	},
}
