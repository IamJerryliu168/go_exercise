package main

import "fmt"

func main() {
	a := [5]int{1,2,3,4,5}
	b := [5]int{1,2,3,4,5}
	c := [5]int{1,2,3}


	fmt.Printf("comapre a == b is %t\n", a == b)
	fmt.Printf("comapre a == c is %t\n", a == c)

	var d [5]int
	d = a
	fmt.Printf("comapre d == a is %t\n", d == a)
}