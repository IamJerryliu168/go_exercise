package main

import "fmt"


func Add(a, b int) (result int) {
	return a + b
}

func Minus(a, b int) (result int) {
	return a - b
}

type FuncType func(int,int) int



func main() {
	result := Add(1, 2)
	fmt.Printf("The first add result is %d\n", result)

	var addfuncTpe FuncType = Add
	result = addfuncTpe(10, 20)
	fmt.Printf("The second add result is %d\n", result)

	result = Minus(2, 1)
    fmt.Printf("The first minus result is %d\n", result)

    var minusFuncType FuncType = Minus
    result = minusFuncType(20, 10)
    fmt.Printf("The second minus result is %d\n", result)

}