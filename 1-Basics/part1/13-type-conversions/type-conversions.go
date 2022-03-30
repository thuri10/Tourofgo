package main

import (
	"fmt"
	"math"
)

func main() {
	//intialize variable x and y
	var x, y int = 3, 4
	//use float type because of decimal point
	var f float64 = math.Sqrt(float64(x*x + y*y)) //sqrt of 9 and 16 = 25
	var z uint = uint(f) //convert float value to uint
	fmt.Println(x, y, z)
}