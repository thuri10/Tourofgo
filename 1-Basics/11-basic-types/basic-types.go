package main

import (
	"fmt"
	"math/cmplx"
)
//declare and initialize variables

var (
	ToBe    bool = false
	//use int unless you have a specific reason to used signed and unsigned types
	MaxInt  uint64 = 1<<64 -1
	z       complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	//boolean type
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	//int type
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	//complex128 type
	fmt.Printf("Type: %T Value: %v\n", z, z)

}