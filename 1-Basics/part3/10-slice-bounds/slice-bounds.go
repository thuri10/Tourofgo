// only gets values of a given slice boundary
package main

import "fmt"

func main() {
	s :=[]int{2, 3, 5, 7, 11, 13}
    //elements bound by index 1 and 4
	s = s[1:4]
	fmt.Println(s)
    //last two elements
	s = s[:2]
	fmt.Println(s)
    //elements from index 1
	s = s[1:]
	fmt.Println(s)
}