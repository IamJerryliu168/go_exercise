package main

import "fmt"


func main() {
	var m1 map[int]string
	fmt.Println("m1=", m1)
	fmt.Println("len(m1)=",len(m1))

	m2:=make(map[int]string)
	fmt.Println("m2=", m2)
	fmt.Println("len(m2)=",len(m2))

	m3:=make(map[int]string,1)	
	fmt.Println("m3=", m3)
	fmt.Println("len(m3)=",len(m3))
	m3[0] = "a"
	m3[1] = "b"
	m3[2] = "c"
    fmt.Println("after add element m3=", m3)
	fmt.Println("after add element len(m3)=",len(m3))

	m4:=map[int]string{1:"java",2:"shell",3:"python",4:"go"}
	fmt.Println("m4=", m4)
	fmt.Println("len(m4)=",len(m4))


	// map traversal
	for key,value := range m4 {
		fmt.Printf("m4's key %d, value %s\n", key,value)
	}
	// to judge the value if in map
	value, ok := m4[1]
	if ok == true {
		fmt.Println("the key in m4 exist!")
		fmt.Println("the value is ", value)
	} else {
		fmt.Println("the key in m4 not exist!")
	}

	// map delete element by key
	delete(m4,4)
	fmt.Println("after delete element m4=", m4)

}