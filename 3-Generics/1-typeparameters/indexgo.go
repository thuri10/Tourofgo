package main

import "fmt"

//Index returns the index of x in s, or -1 if not found

//does not run on local machine
func Index[T comparable] (s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		//we can use == or != constrains here

		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	//Index also works for slice of strings

	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))
}