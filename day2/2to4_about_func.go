package main

import "fmt"



func MyFunc01() {
	fmt.Println("This is a func which no parameters and no return values.")
}


func MyFunc02(a int, str string) {
	fmt.Println("This is a func with standard parameters, but still no return values")
	fmt.Printf("The para a = %d, para str = %s\n", a, str)
}

func MyFunc03(args ...int) {

	for i := 0; i < len(args); i++ {
		fmt.Println("with simple for using, the value is ", args[i])
	}

	for index, value := range args {
		fmt.Printf("with range, the index is %d, the value is %d\n", index, value)
	}
}


func main() {
	MyFunc01()
	MyFunc02(10, "This is the func test.")
	MyFunc03()
	MyFunc03(5,6,7,8,9,10)
}