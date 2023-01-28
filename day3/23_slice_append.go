package main

import "fmt"


func main() {
	s1 := []int{}
	fmt.Println("s1=",s1)
	fmt.Printf("s1 len = %d, cap = %d\n", len(s1), cap(s1))

	s1 = append(s1, 1)
	s1 = append(s1, 2)
	fmt.Println("after append s1=",s1)
	fmt.Printf("after append s1 len = %d, cap = %d\n", len(s1), cap(s1))


	var s2 = []int{1,2,3,4}
	s2 = append(s2,5)
	s2 = append(s2,5)
	s2 = append(s2,5)
	fmt.Println("after append s2=",s2)
	fmt.Printf("after append s2 len = %d, cap = %d\n", len(s2), cap(s2))	

}
