package main

import "fmt"



func MyFunc01(paras ...string) {

	fmt.Println("the paras's length is ", len(paras))

	for index, value := range paras {
		fmt.Printf("The index is %d, the value is %s \n", index, value)
	}
}




func Test(args ...string) {
	// pass all parameters
	MyFunc01(args ...)

	// pass the parameters begin from the third parameter ( include the third parameter)
	MyFunc01(args[3:] ...)

	// pass the parameters begin from the first parameter until the third parameter (not include the third parameter)
	MyFunc01(args[:3] ...)
}




func main() {

	Test("a","b","c","d","e","f","g","h","i")
}