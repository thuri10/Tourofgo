package main

import "fmt"

//the value stored in Pi variable does not change
//constants cannot be declared using type inference
const Pi = 3.142

func main() {
	fmt.Println("Happy", Pi, "Day")
}