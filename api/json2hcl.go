package main

import (
	"log"
	"encoding/json"
	"encoding/gob"
	"fmt"
	"bytes"
	"io"
	//"os/exec"
	"io/ioutil"
	//hcl "github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/printer"
	jsonParser "github.com/hashicorp/hcl/json/parser"
	"net/http"
	//"github.com/go-playground/form"
	//"github.com/gorilla/mux"
)

var err error

//func toHCL(input []byte) error {
//
//	ast, err := jsonParser.Parse([]byte(input))
//	if err != nil {
//		return fmt.Errorf("unable to parse JSON: %s", err)
//	}
//
//	err = printer.Fprint(os.Stdout, ast)
//	if err != nil {
//		return fmt.Errorf("unable to print HCL: %s", err)
//	}
//
//	return nil
//}

//func toHCL(input []byte) error {
//	//input, err := ioutil.ReadFile("input.json")
//	if err != nil {
//		return fmt.Errorf("unable to read from stdin: %s", err)
//	}
//
//	ast, err := jsonParser.Parse([]byte(input))
//	if err != nil {
//		return fmt.Errorf("unable to parse JSON: %s", err)
//	}
//
//	//err = printer.Fprint(os.Stdout, ast)
//	//err = ioutil.WriteFile("out.tf", GetBytes(ast), 0644)
//	if err != nil {
//		return fmt.Errorf("unable to print HCL: %s", err)
//	}
//
//	return nil
//}

func GetBytes(key interface{}) []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(key)
	return buf.Bytes()
}

//var decoder *form.Decoder

func Json2Hcl(w http.ResponseWriter, r *http.Request) {

	/*
	 Client must set header : -H "Content-Type: application/x-www-form-urlencoded"
	*/

	 // 1. Get the datas from http request
	 // 2. parse them
	 // 3. print them to a json file.

	w.Header().Set("Content-Type", "plain/text; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	fmt.Println("Request received.")

	type Data struct {
		key string
		value string
	}

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

	//fmt.Printf("body: %s\n", body)
	//fmt.Printf("d: %s\n", d)

	ast, err := jsonParser.Parse(body)
	if err != nil {
		panic(err)
	}

	log.Println(d)
	if err := printer.Fprint(w, ast); err != nil {
	//if err := json.NewEncoder(w).Encode(ast); err != nil {
		panic(err)
	}
}
