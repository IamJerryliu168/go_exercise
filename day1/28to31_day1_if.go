
package main

import "fmt"

func main() {
  
  // 1
  a:= 10
  if a==10 {
  	fmt.Println("a = 10")
  }	

  // 2
  if b:=20; b == 10 {
  	 fmt.Println("b = 10")
  } else {
  	 fmt.Println("b != 10")
  }

  //3
  var c = 20
  if c == 30 {
    fmt.Println("c = 30")
  } else if c > 30 {
  	fmt.Println("c > 30")
  } else if c < 30 {
  	fmt.Println("c < 30")
  } else {
  	fmt.Println("It is impossible!")
  }


}	