package main

import "fmt"



func main() {
	var a int = 10
	fmt.Printf("the a value is %d\n", a)
	fmt.Printf("the a address is %p\n", &a)


	//Pointer variables
	var p *int
	p = &a
	fmt.Printf("p = %v, &a = %v \n", p, &a)
	*p = 666
	fmt.Printf("a value is %d \n", a)


	// when not assign value to poiner variables, the default value is nil
	// at this time if *p = 777 will throw error
	var q *int    // nil
	fmt.Printf("q = %v\n", q)

	// use new(int) for poiner variables init
	r := new(int)    // new(int) will return the address 
	*r = 777
	fmt.Printf("*r = %v\n", *r)
}