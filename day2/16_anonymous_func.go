package main

import "fmt"


func main() {

	value := 100
	name := "This is the GO!"

	f1 := func() {
		fmt.Println("The value is ", value)
		fmt.Println("The name is ", name)
	}


	type FuncType func()

	var f2 FuncType
	f2 = f1
	f2()

	func () {
		fmt.Printf("The value is %d, the name is %s\n", value, name)
	}()

	max,min := func (a, b int)(max, min int) {
		if a>b {
			max = a
			min = b
		} else {
			max = b
			min = a 
		}
		return
	}(10, 20)
	fmt.Printf("The max value is %d, the min value is %d\n", max, min)

}