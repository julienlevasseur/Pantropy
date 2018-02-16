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

func RouterOSProvTest(route string, method string) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc(route, ProvidersOpenstack).Methods(method)
	return router
}

func TestProvidersOS(t *testing.T) {
	request, _ := http.NewRequest("GET", "/v1/providers/openstack", nil)
	response := httptest.NewRecorder()
	RouterProvTest("/v1/providers/openstack", "GET").ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	if response.Code == http.StatusOK {
		displayStatusOK("ProvidersOpenstack", response.Code)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("ProvidersOpenstack did not panic")
		} else {
			PrintGreen("[OK] ProvidersOpenstack panic correctly")
		}
	}()

	var w http.ResponseWriter
	var r *http.Request
	ProvidersOpenstack(w, r)
}

func RouterResOSTest(route string, method string) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc(route, ResourceOpenstack).Methods(method)
	return router
}

func TestResOS(t *testing.T) {
	request, _ := http.NewRequest("GET", "/v1/providers/openstack/openstack_instance", nil)
	response := httptest.NewRecorder()
	RouterResAWSTest("/v1/providers/openstack/openstack_instance", "GET").ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	if response.Code == http.StatusOK {
		displayStatusOK("ResourceOpenstack", response.Code)
	}

	openstackRes := openstackResources
	if string(reflect.TypeOf(openstackRes).Kind().String()) == "map" {
		displayTypeOK("openstackResources", string(reflect.TypeOf(openstackRes).Kind().String()))
	} else {
		t.Errorf(color.RedString("wrong openstackResources type: got %v want %v"),
			string(reflect.TypeOf(openstackRes).Kind().String()),
			"map",
		)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("ResourceOpenstack did not panic")
		} else {
			PrintGreen("[OK] ResourceOpenstack panic correctly")
		}
	}()

	var w http.ResponseWriter
	var r *http.Request
	ResourceOpenstack(w, r)
}
