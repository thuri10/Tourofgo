package main

import "fmt"

func sum(s []int, c chan int){
	sum :=0
	for _, v := range s {
		sum += v
	}
	//send the value of sum to the channel
	c <- sum
}

func main() {
	s := []int {7, 2, 8, -9, 4, 0}

	//make a channel before using 
	c := make(chan int) 

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c  //receive from  channel c
	fmt.Println(x, y, x+y)
}