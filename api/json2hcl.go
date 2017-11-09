package main

import (
	"log"
	"encoding/json"
	"encoding/gob"
	"fmt"
	"bytes"
	"io"
	"io/ioutil"
	"github.com/hashicorp/hcl/hcl/printer"
	jsonParser "github.com/hashicorp/hcl/json/parser"
	"net/http"

var err error

func Json2Hcl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "plain/text; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	var d Data
	if err := json.Unmarshal(body, &d); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	ast, err := jsonParser.Parse(body)
	if err != nil {
		panic(err)
	}

	log.Println(d)
	if err := printer.Fprint(w, ast); err != nil {
		panic(err)
	}
}
