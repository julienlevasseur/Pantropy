package main

/*
* This file contain the Application resources routes.
 */

import (
	"encoding/json"
	"net/http"
)

func AppRoutes(w http.ResponseWriter, r *http.Request) {

	var infraRoutes [1]string
	infraRoutes[0] = "docker"

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(infraRoutes); err != nil {
		panic(err)
	}
}

func DockerRoutes(w http.ResponseWriter, r *http.Request) {

	var infraRoutes [3]string
	infraRoutes[0] = "images"
	infraRoutes[1] = "build"
	infraRoutes[2] = "push"

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(infraRoutes); err != nil {
		panic(err)
	}
}
