package main

import "fmt"

var a byte

func main() {

	a := 8
	fmt.Printf("1: the a's type is %T \n", a)
	{
		var a float32
		fmt.Printf("2: the a's type is %T \n", a)
	}

	test()
}


func test() {
  fmt.Printf("3: the a's type is %T \n", a)
}