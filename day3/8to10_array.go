package main

import "fmt"

func main() {
	var a1 [5]int
	for i := 0; i< len(a1); i++ {
		a1[i] = i+1
	}

    a2 := [5]{6,7,8,9}

    fmt.Println("a1 = ", a1)
    fmt.Println("a2 = ", a2)

}