package main

import "fmt"

func main() {
	fmt.Println("Counting")
	
	//defer using a loop
	for i:= 0;  i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}