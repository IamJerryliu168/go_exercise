package main

import "fmt"


type Student struct {
	id int
	name string
	age int
	sex byte
	addr string
}


func (tmp Student) setInfoValue() {
	fmt.Println("setInfoValue")
}

func (tmp Student) setInfoPointer() {
	fmt.Println("setInfoPointer")
}

func main() {
	s := Student{1, "student02", 19, 'f', "the hunan provience"}
	s.setInfoPointer()   // in inner, first update s to &s then to invoke
	s.setInfoValue() 
}