package main

import (
	"net/http"
	"encoding/json"
)

func Index(w http.ResponseWriter, r *http.Request) {
	index_routes := []string{"providers"}
	//index_routes := make(map[string][]string)
	//index_routes["result"] = append(index_routes["result"], "providers")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(index_routes); err != nil {
		panic(err)
	}
}

func ProvidersIndex(w http.ResponseWriter, r *http.Request) {
	providers := []Provider{
		Provider{Name: "aws"},
		Provider{Name: "chef"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(providers); err != nil {
		panic(err)
	}
}
