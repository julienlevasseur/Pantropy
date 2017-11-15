package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func RouterProvTest(route string, method string) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc(route, ProvidersAWS).Methods(method)
	return router
}

func TestProvidersAWS(t *testing.T) {
	request, _ := http.NewRequest("GET", "/v1/providers/aws", nil)
	response := httptest.NewRecorder()
	RouterProvTest("/v1/providers/aws", "GET").ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	if response.Code == http.StatusOK {
		displayStatusOK("ProvidersAWS", response.Code)
	}
}

func RouterResAWSTest(route string, method string) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc(route, ResourceAWS).Methods(method)
	return router
}

func TestResAWS(t *testing.T) {
	request, _ := http.NewRequest("GET", "/v1/providers/aws/aws_instance", nil)
	response := httptest.NewRecorder()
	RouterResAWSTest("/v1/providers/aws/aws_instance", "GET").ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	if response.Code == http.StatusOK {
		displayStatusOK("ResourceAWS", response.Code)
	}
}
