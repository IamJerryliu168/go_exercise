package main

import "fmt"


type Person struct {
	name string
	age int
	sex byte
}

func (tmp *Person) PrintInfo() {
	fmt.Printf("person's method ::: name = %s, age = %d, sex = %c\n", tmp.name, tmp.age, tmp.sex)
}


// the student inherit the person, so on Student you can invoke method PrintInfo()
type Student struct {
	Person // Anonymous field
	id int
	addr string
}

func (tmp *Student) PrintInfo() {
	fmt.Println("studnet's method ::: tmp =", tmp)
}

func main() {
	s := Student{Person{"mike",18,'m'}, 8, "the address"}
	s. PrintInfo()  // this will invoking Printnfo() method on Student
	s.Person.PrintInfo() // this will invoking Printnfo() method on Person
}