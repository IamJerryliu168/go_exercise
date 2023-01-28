package main

import "fmt"

//subset
type Humaner interface {
	sayHi()
}

//superset
type Personer interface {
	Humaner
	sing(name string)
}

type Student struct {
	name string
	id int
}

func (tmp *Student) sayHi() {
	fmt.Printf("Student[%s, %d] sayHi \n", tmp.name, tmp.id)
}

func (tmp *Student) sing(name string) {
	fmt.Printf("Student[%s, %d] sing %s \n", tmp.name, tmp.id, name)
}

func main() {
	var h Humaner
	var p Personer

    //superset can transfer to subset
	h = P
	// subset not allow transfer to superset
	//p = h

}