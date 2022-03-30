package main

import "fmt"

func main() {
	//it is printed after execution of Hello returns
	defer fmt.Println("World")

	fmt.Println("Hello")
}