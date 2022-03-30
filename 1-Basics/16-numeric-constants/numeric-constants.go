package main

import "fmt"

const (
	// create a huge number by shifting a 1 bit left 100 places
	Big = 1 <<100

	Small = Big >> 99
)

func needInt(x int) int { return x*10 +1}
func needFloat(x float64) float64 {
	return x* 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Big))
	fmt.Println(needFloat(Small))
	//fmt.Println(needInt(Big)) overflows int type can hold
}