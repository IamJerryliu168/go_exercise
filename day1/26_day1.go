package main

import "fmt"

func main() {

	type boolean bool
	var flag boolean = true

	fmt.Printf("flag's type is %T\n", flag)

	type (
		long int64
		char byte
		String string
	)

	var l long = 1000
	var c char = 'a'
	var str String = "This is the GO!"

	fmt.Printf("l's type is %T, c's type is %T, str's type is %T", l, c, str)

}	