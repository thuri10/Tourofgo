package main

import (
	"fmt"
	"math/rand"
	
)

//entry point of our program
func main() {
	//use a different seed value
	fmt.Println("My Favorite number is ", rand.Intn(12))
}