package main

import (
	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func RouterChefProvTest(route string, method string) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc(route, ProvidersChef).Methods(method)
	return router
}

func TestProvidersChef(t *testing.T) {
	request, _ := http.NewRequest("GET", "/v1/providers/chef", nil)
	response := httptest.NewRecorder()
	RouterProvTest("/v1/providers/chef", "GET").ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	if response.Code == http.StatusOK {
		displayStatusOK("ProvidersChef", response.Code)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("ProvidersChef did not panic")
		} else {
			PrintGreen("[OK] ProvidersChef panic correctly")
		}
	}()

	var w http.ResponseWriter
	var r *http.Request
	ProvidersChef(w, r)
}

func RouterResChefTest(route string, method string) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc(route, ResourceChef).Methods(method)
	return router
}

func TestResChef(t *testing.T) {
	request, _ := http.NewRequest("GET", "/v1/providers/chef/chef_instance", nil)
	response := httptest.NewRecorder()
	RouterResAWSTest("/v1/providers/chef/chef_instance", "GET").ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	if response.Code == http.StatusOK {
		displayStatusOK("ResourceChef", response.Code)
	}

	chefRes := chefResources
	if string(reflect.TypeOf(chefRes).Kind().String()) == "map" {
		displayTypeOK("chefResources", string(reflect.TypeOf(chefRes).Kind().String()))
	} else {
		t.Errorf(color.RedString("wrong chefResources type: got %v want %v"),
			string(reflect.TypeOf(chefRes).Kind().String()),
			"map",
		)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("ResourceChef did not panic")
		} else {
			PrintGreen("[OK] ResourceChef panic correctly")
		}
	}()

	var w http.ResponseWriter
	var r *http.Request
	ResourceChef(w, r)
}
