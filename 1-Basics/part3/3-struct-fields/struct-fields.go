package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	//modify X field with value 4
	v.X = 4
	fmt.Println(v.X)
}