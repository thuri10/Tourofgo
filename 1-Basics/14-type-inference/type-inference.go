package main

import "fmt"

func main() {
	//type inference 
	//one does not need to specify the type of variable when using inference method
	// v variable can be a string, float64 or complex
	v := 42.0
	fmt.Printf("v is of type %T\n", v)
}