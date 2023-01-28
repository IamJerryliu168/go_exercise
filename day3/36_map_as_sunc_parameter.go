package main

import "fmt"


func test(m map[int]string) {
	delete(m,3)
}

func main() {
	m := map[int]string{1:"java",2:"shell",3:"python",4:"go"}
	test(m)
	fmt.Println("m=",m)
}
