package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i    //reference value stored at i
	fmt.Println(*p) //print the value at i

	*p = 21; //update the value at memory address pointed by p
	fmt.Println(i)
	
	p = &j // point to j
	*p = *p / 37 // divide j by 37
	fmt.Println(j) //j is value at memory *p

}