package main

import "fmt"


func main() {

	var str1 string
	fmt.Println("str1=",str1)
	str1 = "This is the GO"
	fmt.Println("str1 length is", len(str1))

    str2 := "This is the JAVA"
    fmt.Printf("str2 type is %T", str2)

}