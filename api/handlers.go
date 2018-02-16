package main

import (
	"encoding/json"
	"net/http"
)

// Index : Define the index fuction
func Index(w http.ResponseWriter, r *http.Request) {
	indexRoutes := []string{"providers"}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(indexRoutes); err != nil {
		panic(err)
	}
}

// ProvidersIndex : Define the providers index fuction
func ProvidersIndex(w http.ResponseWriter, r *http.Request) {

	var providers [3]string
	providers[0] = "aws"
	providers[1] = "chef"
	providers[2] = "openstack"

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(providers); err != nil {
		panic(err)
	}
}
