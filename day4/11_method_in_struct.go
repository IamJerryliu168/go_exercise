package main

import "fmt"


type Student struct {
	id int
	name string
	age int
	sex byte
	addr string
}


func (tmp Student) printStudentInfo() {
	fmt.Println("tmp = ", tmp)
}

func (p *Student) setStudentInfo() {
	(*p).id = 8
	(*p).name = "fafa"
	(*p).age = 20
	(*p).sex = 'm'
	(*p).addr = "ABCDEWDF"

	fmt.Println("*p = ", *p)

    fmt.Printf("*p = %+v\n", *p)
}

func main() {

	s1 := Student{1, "yaofa", 18, 'f',"KKKKKKK"}
	s1.printStudentInfo()
	s1.setStudentInfo()
}