package main

import (
	//"reflect"
	"net/http"
	"net/http/httptest"
	"testing"
	//"github.com/gorilla/mux"
	//"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
)

func TestJSON2Hcl(t *testing.T) {
	request, _ := http.NewRequest("POST", "/v1/json2hcl", nil)
	response := httptest.NewRecorder()
	RouterResAWSTest("/v1/json2hcl", "POST").ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	if response.Code == http.StatusOK {
		displayStatusOK("JSON2Hcl", response.Code)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("JSON2Hcl did not panic")
		} else {
			PrintGreen("[OK] JSON2Hcl panic correctly")
		}
	}()

	var w http.ResponseWriter
	var r *http.Request
	JSON2Hcl(w, r)
}
