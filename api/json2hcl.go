package main

import (
	"io"
	"io/ioutil"
	"github.com/hashicorp/hcl/hcl/printer"
	jsonParser "github.com/hashicorp/hcl/json/parser"
	"net/http"
)

var err error

// JSON2Hcl convert JSON to HCL
func JSON2Hcl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "plain/text; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	ast, err := jsonParser.Parse(body)
	if err != nil {
		panic(err)
	}

	if err := printer.Fprint(w, ast); err != nil {
		panic(err)
	}
}
