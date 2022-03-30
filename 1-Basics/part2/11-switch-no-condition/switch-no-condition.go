package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	//switch can be used as if else statement condition
	switch {
		case t.Hour() < 12:
			fmt.Println("Good Morrning")
		case t.Hour() < 17:
			fmt.Println("Good afternoon")
		default:
			//switch case should have a default case condition
			fmt.Println("Good evening")
	}
}