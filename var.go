package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

//inferred global declaration
var (
	name   = os.Getenv("LOGNAME")
	course = "Go Fundamentals"
	module = "4" //now a string
	clip   = 2
)

//runtime reflection
func main() {

	courseComplete := false

	fmt.Println("Name and course are set to", name, "and", course, ".")
	fmt.Println("Module and clip are set to", module, "and", clip, ".")
	fmt.Println("Name is of type", reflect.TypeOf(name))
	fmt.Println("Module is of type", reflect.TypeOf(module))
	fmt.Println("Course completed:", courseComplete)
	//total := module + clip
	//fmt.Println("Module + clip equals", total)

	// error check
	iModule, err := strconv.Atoi(module)
	if err == nil {
		total := iModule + clip //local variable
		fmt.Println("Module + clip equals", total)
	}

	//pointer variables
	fmt.Println("Memory address of *course* is", &course)

	var ptrVar *string = &course
	fmt.Println("Pointing course variable at address,", ptrVar, "Which holds the value", *ptrVar)
}
