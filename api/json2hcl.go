package main

import (
	"log"
	"encoding/json"
	"encoding/gob"
	"bytes"
	"fmt"
	"os"
	//"os/exec"
	//"io/ioutil"
	//hcl "github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/printer"
	jsonParser "github.com/hashicorp/hcl/json/parser"
	"net/http"
	"github.com/go-playground/form"
	//"github.com/gorilla/mux"
)

var err error

func toHCL(input []byte) error {

	ast, err := jsonParser.Parse([]byte(input))
	if err != nil {
		return fmt.Errorf("unable to parse JSON: %s", err)
	}

	err = printer.Fprint(os.Stdout, ast)
	if err != nil {
		return fmt.Errorf("unable to print HCL: %s", err)
	}

	return nil
}

func GetBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

var decoder *form.Decoder

func Json2Hcl(w http.ResponseWriter, r *http.Request) {

	/*
	 Client must set header : -H "Content-Type: application/x-www-form-urlencoded"
	*/
	w.Header().Set("Content-Type", "plain/text; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	decoder = form.NewDecoder()

	r.ParseForm() // Parse the request body
	values := r.Form

	var input []byte

	log.Println(values)
	
	err := decoder.Decode(&input, values)
	if err != nil {
		log.Println(err)
	}

	var i []byte
	log.Println(input)
	i, err = GetBytes(input)
	log.Println(i)
	
	if err := json.NewEncoder(w).Encode(toHCL(i)); err != nil {
		panic(err)
	}
}
