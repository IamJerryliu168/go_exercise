package main

import "fmt"


type Student struct {
	id int
	name string
	age int
	sex byte
	addr string
}

func notModifyInMain(s Student) {
	s.name = "ABC"
	fmt.Println("in notModifyInMain:::  s=", s)
}

func modifyInMain(p *Student) {
    (*p).name = "EFG"
}

func main() {
   s := Student{1,"zhangsan",20,'m',"Beijing"}	
  // test value pass
  notModifyInMain(s)
  fmt.Println("in main:::  s=", s)

  // test address pass
  modifyInMain(&s)
  fmt.Println("in main2:::  s=", s)


}	