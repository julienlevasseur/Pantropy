package main

import (
	"net/http"
	"encoding/json"
)

// Define the index fuction
func Index(w http.ResponseWriter, r *http.Request) {
	index_routes := []string{"providers"}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(index_routes); err != nil {
		panic(err)
	}
}

// Define the providers index fuction
func ProvidersIndex(w http.ResponseWriter, r *http.Request) {
	// Actually, the only two supported providers are : aws, chef
	providers := []Provider{
		Provider{Name: "aws"},
		Provider{Name: "chef"},
		Provider{Name: "openstack"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(providers); err != nil {
		panic(err)
	}
}
