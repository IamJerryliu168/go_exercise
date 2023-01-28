package main

import "fmt"


type Student struct {
	name string
	id int
}


func xxx(arg ...interface{}) {

}


func main() {

	var i interface{} = 1
	fmt.Println("i =", i)
	var str interface{} = "the null interface"
	fmt.Println("str =", str)
	var student interface{} = Student{"student01", 1}
	fmt.Printf("the student %+v\n", student)

	var v1 interface{} = &str
	fmt.Println("v1 =", v1)
	var v2 interface{} = struct{X int}{1}
	fmt.Println("v2 =", v2)
	var v3 interface{} = &struct{X int}{1}
	fmt.Println("v3 =", v3)

}