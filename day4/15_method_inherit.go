package main

import "fmt"


type Person struct {
	name string
	age int
	sex byte
}

func (tmp * Person) PrintInfo() {
	fmt.Printf("name = %s, age = %d, sex = %c\n", tmp.name, tmp.age, tmp.sex)
}


// the student inherit the person, so on Student you can invoke method PrintInfo()
type Student struct {
	Person // Anonymous field
	id int
	addr string
}


func main() {
	s := Student{Person{"mike", 20, 'f'}, 1, "abbb"}
    
    //student s invoke method PrintInfo() which belong to Person 
	s.PrintInfo()
}