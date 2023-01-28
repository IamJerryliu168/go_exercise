package main

import "fmt"


func swap(p1, p2 *int){
	*p1,*p2 = *p2,*p1
}

func main() {
	a:=10
	b:=20
	swap(&a,&b)
	fmt.Printf("main::: a=%d,b=%d\n", a, b)
}