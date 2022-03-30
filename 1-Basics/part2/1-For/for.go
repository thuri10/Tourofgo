package main

import "fmt"

func main() {
	sum := 0
	//for does not have brackets compared to c/c++ or JS
	for i := 0; i < 10; i++ {
		sum +=i
	}
	fmt.Println(sum)
}