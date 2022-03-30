package main

import "fmt"

//swap two strings function
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	//infer the type of variables a and b using :=
	a, b := swap("hello", "world")
	//a -> world
	//b -> hello
	fmt.Println(a, b)
}