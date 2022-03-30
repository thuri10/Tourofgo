package  main

import "fmt"

//function to do add calculation
//the return value is an interger
//pass tow variables of type interger
func add( x int, y int) int {
	return x + y
}

func main() {
	//call add function and pass tow arguments
	fmt.Println(add(42, 13))
}