package main

import (
	"fmt"
	"github.com/fatih/color"
)

func displayStatusOK(route string, status int){
	fmt.Printf(color.GreenString("%v%v%v%v\n"), "[OK] ", route, " return: ", status)
}

func displayBodyOK(route string, body string){
	fmt.Printf(color.GreenString("%v%v%v%v\n"), "[OK] ", route, " body: ", body)
}

func displayTypeOK(obj string, obj_type string){
	fmt.Printf(color.GreenString("%v%v%v%v\n"), "[OK] ", obj, " type: ", obj_type)
}

func PrintGreen(content string){
	fmt.Printf(color.GreenString("%v\n"), content)
}
