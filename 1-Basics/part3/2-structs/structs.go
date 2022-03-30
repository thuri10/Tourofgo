package main

import "fmt"

//put type because golang is typed language
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}