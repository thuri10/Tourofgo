package main

import(
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0  //z is a float64

	// for loop statement
	for i := 0;  i < 10; i++ {
		
		//NEWTONS METHOD FOR sqrt
		z -= (z*z -x) / (2 *z)
		
		//fmt.Println(z)  // repeated until z value is constant
	}
	return z
}

func main() {
	fmt.Println(Sqrt(1000))
}