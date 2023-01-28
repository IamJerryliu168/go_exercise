package main

import "fmt"


func main() {

	var a = [5]int{1,2,3,4,5}
	fmt.Println("arryy a is ", a)
	s1 := a[0:3:5]
	fmt.Println("slcie s1 is ", s1)

	s2:=[]int{}  // this is a slice
    fmt.Println("slcie s2 is ", s2)
    fmt.Printf("s2 length is %d, capacity is %d \n", len(s2), cap(s2))
    s2 = append(s2, 1,2,3,4)
    fmt.Printf("s2 after append :: length is %d, capacity is %d \n", len(s2), cap(s2))

    // make(slice_type, len, cap)
    s3:= make([]int, 5)
    fmt.Printf("s3 length is %d, capacity is %d \n", len(s3), cap(s3))

    s4:= make([]int, 8, 10)
    fmt.Printf("s4 length is %d, capacity is %d \n", len(s4), cap(s4))


}
