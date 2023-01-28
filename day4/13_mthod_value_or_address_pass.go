package main

import "fmt"


type Student struct {
	id int
	name string
	age int
	sex byte
	addr string
}


func (tmp Student) setStudentInfoByValue() {
	tmp.id = 8
	fmt.Println("tmp = ", tmp)
	fmt.Printf("tmp address is %p\n", &tmp)
}


func (tmp *Student) setStudentInfoByPointer() {
	tmp.id = 9
	fmt.Println("*tmp = ", *tmp)
	fmt.Printf("tmp address is %p\n", tmp)
}

func main() {
	var s = Student{1, "student01", 18, 'f', "QQQQQQ"}
	fmt.Printf("the s address is %p\n", &s)
	s.setStudentInfoByValue()
	s.setStudentInfoByPointer()
}