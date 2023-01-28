package main

import "fmt"


type Student struct {
	id int
	name string
	age int
	sex byte
	addr string
}

func main() {
    // struct's compare
	s1:= Student{1,"zhangsan",18,'m',"Hunan"}
	s2:= Student{1,"zhangsan",18,'m',"Hunan"}
	s3:= Student{1,"zhangsan",18,'m',"hunan"}

	fmt.Println("s1 == s2", s1 == s2)
	fmt.Println("s1 == s3", s1 == s3)

	//compare's assign
	var s4 Student
	s4 = s3
	fmt.Println("s4 = ", s4)

}
