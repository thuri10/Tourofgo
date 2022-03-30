package main

//use this style for importing packages
import (
	"fmt"
	"math"
)

func main() {
	//add \n because of use of Printf macro
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}