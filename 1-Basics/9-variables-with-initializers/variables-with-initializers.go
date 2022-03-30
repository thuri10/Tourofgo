package main

import "fmt"

// variables i and j are int with value 1 and 2 respectively
var i, j int = 1, 2

func main() {
	//if variable is initialized, the type can be omittedog
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}