package main

import (
	"errors"
	"fmt"
)

func div(a int, b int) (result int, err error) {
	err = nil

	if b == 0 {
		err = errors.New("the denominator is zero!")
	} else {
		result = a / b
	}
	return
}

func main() {
	error1 := fmt.Errorf("the first error!")
	error2 := errors.New("the second error!")
	fmt.Printf("the first error: %s,\n the second error: %s \n", error1, error2)

	myResult, err := div(10, 0)
	if err != nil {
		fmt.Printf("error= %s \n", err)
	} else {
		fmt.Printf("the result is %d \n", myResult)
	}
}
