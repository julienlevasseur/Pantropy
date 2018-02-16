package main

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {
	router := NewRouter()
	assert.Equal(t, "*mux.Router", string(reflect.TypeOf(router).String()),
		"main did not return a *mux.Router type",
	)
}
