package main

import "fmt"


func Add(a,b int) int {
	return a + b
}

func Minus(a,b int) int {
	return a - b 
}

type FuncType func(a, b int) int 

func test(a, b int, funcType FuncType) int {
	return funcType(a, b)
}

func main() {

	result := test(1, 2, Add)
	fmt.Printf("the add result is %d\n", result)

	result = test(20, 10, Minus)
	fmt.Printf("the minus result is %d\n", result)
}