package main

import (
	"github.com/fatih/color"
	"reflect"
	"testing"
)

func TestRouter(t *testing.T) {
	router := NewRouter()
	if string(reflect.TypeOf(router).String()) == "*mux.Router" {
		displayTypeOK("router", string(reflect.TypeOf(router).String()))
	} else {
		t.Errorf(color.RedString("wrong router type: got %v want %v"),
			reflect.TypeOf(router).String(),
			"*mux.Router",
		)
	}
}
