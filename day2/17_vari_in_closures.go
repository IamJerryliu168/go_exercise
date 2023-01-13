
package main

import "fmt"

func main() {

   age := 30
   name := "GO"

   func() {
   	  age = 40
   	  name = "Java"

   	  fmt.Printf("The inner age is %d, name is %s \n", age, name)
   }()

   fmt.Printf("The outer age is %d, name is %s \n", age, name)

}