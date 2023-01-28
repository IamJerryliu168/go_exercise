package main

import "fmt"

func main() {
	s1:=[]int{8,9}
	s2:=make([]int,5)
	copy(s2, s1)

	fmt.Println("s2=",s2)
}	