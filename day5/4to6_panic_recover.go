package main

import "fmt"

func a() {
	fmt.Println("aaaaaa")
}

func b(i int) {
	// 1
	//fmt.Println("bbbbbbb")
	//panic("this is a panic error")

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	var array [5]int
	array[i] = 100
}

func c() {
	fmt.Println("ccccccccc")
}

func main() {
	a()
	b(20)
	c()

}
