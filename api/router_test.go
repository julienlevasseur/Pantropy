package main

import (
	"reflect"
	"testing"
	"github.com/fatih/color"
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
