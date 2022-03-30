//pointers to structs
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}

	p := &v  //reference memory of v
	p.X = 1e9 // modify X field
	fmt.Println(v)
}