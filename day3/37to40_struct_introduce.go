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
    // struct init method 1, new method will return a address of Studnet
	s1 := new(Student)
	s1.id = 1
	s1.name = "zhagsan"
	s1.age = 12
	s1.sex = 'm'
	s1.addr = "hunan"
	fmt.Println("s1 = ", s1)
	fmt.Println("*s1 =", *s1)

	// struct init method 2
	s2 := Student{2,"wangwu",13,'f',"shanghai"}
	fmt.Println("s2 = ", s2)

	// struct init method 3
	s3 := Student{id:3,name:"lisi",sex:'m'}
	fmt.Println("s3 = ", s3)

	// struct pointer variable
	var p *Student
	p = &s2
	(*p).name = "liulu"
	(*p).sex = 'm'
    fmt.Println("*p = ", *p)

}