package main

import (
	"fmt"
	"strings"
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/fatih/color"
)

func displayStatusOK(route string, status int){
	fmt.Printf(color.GreenString("%v%v%v%v\n"), "[OK] ", route, " return: ", status)
}

func displayBodyOK(route string, body string){
	fmt.Printf(color.GreenString("%v%v%v%v\n"), "[OK] ", route, " body: ", body)
}

func TestIndexRoute(t *testing.T) {
	route := "/v1"
	req, err := http.NewRequest("GET", route, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Index)
	handler.ServeHTTP(rr, req)


	if status := rr.Code; status != http.StatusOK {
		t.Errorf(color.RedString("index route returned wrong status code: got %v want %v"),
			status, http.StatusOK)
	} else {
		//fmt.Printf(color.GreenString("%v%v%v%v\n"), "[OK]", route, " return: ", status)
		displayStatusOK(route, status)
	}

	expected := "[\"providers\"]"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf(color.RedString("handler returned unexpected body: got %v want %v"),
			rr.Body.String(), expected)
	} else {
		displayBodyOK(route, rr.Body.String())
	}
}

func TestProvidersRoute(t *testing.T) {
	route := "/v1/providers"
	req, err := http.NewRequest("GET", route, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Index)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf(color.RedString("%v returned wrong status code: got %v want %v"),
			route, status, http.StatusOK)
	} else {
		displayStatusOK(route, status)
	}
}

func TestAWSRoute(t *testing.T) {
	route := "/v1/providers/aws"
	req, err := http.NewRequest("GET", route, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Index)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf(color.RedString("%v returned wrong status code: got %v want %v"),
			route, status, http.StatusOK)
	} else {
		displayStatusOK(route, status)
	}
}

func TestChefRoute(t *testing.T) {
	route := "/v1/providers/chef"
	req, err := http.NewRequest("GET", route, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Index)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf(color.RedString("%v returned wrong status code: got %v want %v"),
			route, status, http.StatusOK)
	} else {
		displayStatusOK(route, status)
	}
}

func TestOpenStackRoute(t *testing.T) {
	route := "/v1/providers/openstack"
	req, err := http.NewRequest("GET", route, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Index)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf(color.RedString("%s returned wrong status code: got %v want %v"),
			route, status, http.StatusOK)
	} else {
		fmt.Printf(color.GreenString("%v%v\n"), "[OK] 'openstack' return: ", status)
	}
}
