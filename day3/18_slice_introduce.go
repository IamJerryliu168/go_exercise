package main

import "fmt"


func main() {
	array := [...]int{1,2,3,4,5}
	// slice 1
	s1 := array[0:3:5]
	fmt.Println("s1 = ",s1)
	fmt.Println("len(s1) = ", len(s1))   // length
	fmt.Println("cap(s1) = ", cap(s1))   // capcity

	s1 = array[1:4:5]
	fmt.Println("s1 = ",s1)
	fmt.Println("len(s1) = ", len(s1))   // length
	fmt.Println("cap(s1) = ", cap(s1))   // capcity

}