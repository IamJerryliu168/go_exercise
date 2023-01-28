package main

import "fmt"


func main() {

	var a = []int{0,1,2,3,4,5,6,7,8,9}
    // the new slice
	s1 :=  a[2:5]
	s1[1] = 666
	fmt.Println("s1=",s1)
	fmt.Println("a=",a)

	//the new slice again
	s2 := s1[2:7]
	s2[2] = 777
	fmt.Println("s2=",s2)
	fmt.Println("a=",a)	 


}
