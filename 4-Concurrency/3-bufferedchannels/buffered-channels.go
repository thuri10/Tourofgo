package main

import "fmt"

//buffere channels are where the buffer length is passed as second argument of the make
func main() {
	//initialize a channel
	ch :=make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}