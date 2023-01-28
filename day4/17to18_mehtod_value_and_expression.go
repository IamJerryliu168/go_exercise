package main

import "fmt"


type Person struct {
	name string
	age int
	sex byte
}

func (tmp *Person) PrintInfoPointer() {
	fmt.Printf("%p, %v\n", tmp, tmp)
}

func (tmp Person) PrintInfovalue() {
	fmt.Printf("%p, %v\n", &tmp, tmp)
}



func main() {
	p:= Person{"wangwu", 10, 'm'}
	p.PrintInfoPointer()
	
	// mehtod vlue , hide the method receiver
	pFunc1 := p.PrintInfoPointer
	pFunc1()
	pFunc2 := p.PrintInfovalue
	pFunc2()

	// method expression
	f := (*Person).PrintInfoPointer
	f(&p)  //  explicitly pass the receiver  ===> p.PrintInfoPointer()

	f2 := (Person).PrintInfovalue
	f2(p)
}