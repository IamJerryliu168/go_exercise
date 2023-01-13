package main

import "fmt"



func cumulative(i int) int {
	if i == 1 {
		return 1
	}
	return i + cumulative(i-1)
}

func main() {
	total := cumulative(100)
	fmt.Printf("tatal value is %d\n", total)
}
