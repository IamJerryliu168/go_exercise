package main

import "fmt"


func test(i int) {
	if i==1 {
		fmt.Println("i= 1")
		return 
	}
	fmt.Println("i=", i)
	test(i - 1)
}


func main() {
	fmt.Println("enter into main")
	test(10)
	fmt.Println("end main")
}
