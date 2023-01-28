package main

import "fmt"


type studentAddress string
type workerAddress string

type Person struct {
	id int
	name string

}


type Student struct {
	Person            // struct data type anoynamous field
	age int
	sex byte
	studentAddress
}

type Worker struct {
	*Person
	int           // basic data type anoynamous field
	sex byte
	workerAddress
}


func main() {
	s1 := Student{Person{1,"student1"},18,'m',"Hunan"}
	fmt.Println("s1 = ", s1)
	fmt.Printf("s1 = %+v\n", s1)
	fmt.Println(s1.Person, s1.age, s1.sex, s1.studentAddress)

	s2 := Student{Person: Person{name:"student2"}, age:19}
	fmt.Printf("s2 = %+v\n", s2)



	w1 := Worker{&Person{3,"worker1"}, 40, 'f', "ABCDEFG"}
	fmt.Printf("w1 = %+v\n", w1)
	fmt.Println(w1.id, w1.name, w1.int, w1.sex, w1.workerAddress)

	w1.Person = new(Person)
	w1.id = 4
	w1.name = "worker02"
	w1.int = 41
	w1.sex = 'm'
	w1.workerAddress = "Hunan"
	fmt.Println(" updated ::: Person", w1.Person)
	fmt.Printf(" updated ::: w1 = %+v\n", w1)


}