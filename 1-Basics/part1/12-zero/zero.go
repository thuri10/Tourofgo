package main

import "fmt"

func main() {
	//shows the default values of variable types when not initialized
	var i int   //print 0
	var f float64 //print 0
	var b bool // print false
	var s string // prints null string

	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}