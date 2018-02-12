package main

/*
* This file contain the Infrastructure resources routes.
*/

import (
    "net/http"
    "encoding/json"
)

func InfraRoutes(w http.ResponseWriter, r *http.Request) {

    var infraRoutes [1]string
    infraRoutes[0] = "providers"

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(infraRoutes); err != nil {
        panic(err)
    }
}