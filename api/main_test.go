package main

import (
	"reflect"
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestMain(t *testing.T) {
	router := NewRouter()
	assert.Equal(t, "*mux.Router", string(reflect.TypeOf(router).String()),
		"main did not return a *mux.Router type",
	)
}
