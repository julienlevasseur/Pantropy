package main

import (
	"github.com/fatih/color"
	"reflect"
	"testing"
)

func TestProviders(t *testing.T) {
	providers := Provider{}
	if string(reflect.TypeOf(providers).Kind().String()) == "struct" {
		displayTypeOK("providers", string(reflect.TypeOf(providers).Kind().String()))
	} else {
		t.Errorf(color.RedString("wrong providers type: got %v want %v"),
			string(reflect.TypeOf(providers).Kind().String()),
			"struct",
		)
	}
}
