package main

import "fmt"


func modifyOrigin(p *[5]int) {
	(*p)[0] = 1000
}

func notModifyOrigin(a [5]int){
	a[0] = 1000
}

func main() {
	var a  =  [5]int{1,2,3,4,5}
	notModifyOrigin(a)
	fmt.Println("element after notModifyOrigin :: a = ", a)
	modifyOrigin(&a)
	fmt.Println("element after modifyOrigin :: a = ", a)
}
