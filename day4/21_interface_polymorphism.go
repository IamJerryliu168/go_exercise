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

func WhoSayHi(h Humaner) {
	h.sayHi()
}

func main() {
	s := Student{"student01", 1}
	t := Teacher{"benjing", "group1"}
	var str myStr = "Hello Mike"

	WhoSayHi(&s)
	WhoSayHi(&t)
	WhoSayHi(&str)

	humans := make([]Humaner, 3)
	humans[0] = &s 
	humans[1] = &t
	humans[2] = &str
	fmt.Printf("the humans is %+v\n", humans)

	for _, h := range humans {
		h.sayHi()
	}


}