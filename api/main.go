package main

import (
    "log"
    "net/http"
)

// main function that instantiate a router for the REST routes handling
func main() {

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
