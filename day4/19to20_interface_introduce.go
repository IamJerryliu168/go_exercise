package main

import "fmt"


type Humaner interface {
	sayHi()
}

type Student struct {
	name string
	id int
}

func (tmp *Student) sayHi() {
	fmt.Printf("Student[%s, %d] sayHi \n", tmp.name, tmp.id)
}

type Teacher struct {
	addr string
	group string
}

func (tmp *Teacher) sayHi() {
	fmt.Printf("Teacher[%s, %s] sayHi \n", tmp.addr, tmp.group)
}

type myStr string

func (tmp *myStr) sayHi() {
	fmt.Printf("myStr is %+v\n", *tmp)
}


func main() {
    var i Humaner

	s := Student{"student01", 1}
	s.sayHi()
	s2 := Student{"student02", 2}
	i = &s2
	i.sayHi()

	t := Teacher{"hunan", "teacher group"}
	t.sayHi()
	t2 := Teacher{"hunan02", "teacher group 02"}
	i = &t2
	i.sayHi()

	var m myStr
	m = "this is the myStr type"
	m.sayHi()
	var m2 myStr
	m2 = "this is the 02 mystr type"
	i = &m2
	i.sayHi()

}