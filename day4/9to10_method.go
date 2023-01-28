package main

import "fmt"


type number int


func add(a, b int) (result int) {
	return a + b
}


func (reciver number) add(b number) (result number) {
	return reciver + b
}

func main() {
	var a,b = 1,2
	var c number = 1

	r1 := add(a,b)
	fmt.Println("r1 = ", r1)

    var d number
    d = 2
	r2 := c.add(d)
	fmt.Println("r2 = ", r2)
}
