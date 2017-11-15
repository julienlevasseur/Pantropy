package main

import (
	"reflect"
	"testing"
	"github.com/fatih/color"
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
