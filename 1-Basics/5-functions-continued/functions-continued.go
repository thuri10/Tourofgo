package main

import "fmt"

//use x,y int instead of x int, y int
func add(x,y int) int {
	return x + y
}

func main() {
	//call our function add
	fmt.Println(add(45,13))

}