package main

import (
	"reflect"
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/gorilla/mux"
	"github.com/fatih/color"
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

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("ProvidersAWS did not panic")
		} else {
			PrintGreen("[OK] ProvidersAWS panic correctly")
		}
	}()

	var w http.ResponseWriter
	var r *http.Request
	ProvidersAWS(w, r)
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

	awsRes := awsResources
	if string(reflect.TypeOf(awsRes).Kind().String()) == "map" {
		displayTypeOK("awsResources", string(reflect.TypeOf(awsRes).Kind().String()))
	} else {
		t.Errorf(color.RedString("wrong awsResources type: got %v want %v"),
			string(reflect.TypeOf(awsRes).Kind().String()),
			"map",
		)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("ResourceAWS did not panic")
		} else {
			PrintGreen("[OK] ResourceAWS panic correctly")
		}
	}()

	var w http.ResponseWriter
	var r *http.Request
	ResourceAWS(w, r)
}
