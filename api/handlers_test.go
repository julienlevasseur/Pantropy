package main

import (
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RouterIndexTest(route string, method string) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc(route, Index).Methods(method)
	return router
}

func TestHandlerIndex(t *testing.T) {
	request, _ := http.NewRequest("GET", "/v1", nil)
	response := httptest.NewRecorder()
	RouterIndexTest("/v1", "GET").ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	if response.Code == http.StatusOK {
		displayStatusOK("Index", response.Code)
	}
}

func RouterProvIndexTest(route string, method string) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc(route, ProvidersIndex).Methods(method)
	return router
}

func TestHandlerProvIndex(t *testing.T) {
	request, _ := http.NewRequest("GET", "/v1/providers", nil)
	response := httptest.NewRecorder()
	RouterProvIndexTest("/v1/providers", "GET").ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	if response.Code == http.StatusOK {
		displayStatusOK("ProvidersIndex", response.Code)
	}
}
