package main

import (
	"net/http"
	"encoding/json"
)

// Index: Define the index fuction
func Index(w http.ResponseWriter, r *http.Request) {
	indexRoutes := []string{"providers"}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(indexRoutes); err != nil {
		panic(err)
	}
}

// ProvidersIndex: Define the providers index fuction
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
