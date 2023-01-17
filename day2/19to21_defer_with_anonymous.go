package main

import "fmt"


// defer is LIFO 
func main() {

	a := 10
	b := 20

	defer func (a,b int) {
		fmt.Printf("The a is %d, the b is %d in defer.\n", a, b)
	}(a, b)

	a = 100
	b = 200
	fmt.Printf("The a is %d, the b is %d not in defer.\n", a, b)

	fmt.Printf("This is the last defer sentence in main!\n")
}
