package main

import "fmt"

func main() {
  
  //1
  var floor int
  fmt.Println("Pls enter the floor:")
  fmt.Scan(&floor)

  switch floor {
  case 1:
  	fmt.Println("The first floor.")
  	//break
  	//fallthrough
  case 2:
    fmt.Println("The sceond floor.")
  case 3:
    fmt.Println("The third floor.")
  default:
    fmt.Println("The floor > 3.")    	
  }

  //2
  switch num:=5; num {
  case 20 :
  	fmt.Println("The num > 10.")
  case 10:
  	fmt.Println("The num == 10.")
  case 1,2,3,4,5,6,7,8,9:
  	fmt.Println("The num < 10.")
  default:
  	fmt.Println("The num is not 20, but > 10.")
  }

  //3	
  var count = 12
  switch{
  case count > 12:
  	fmt.Println("The count is > 12.")
  case count == 12:
    fmt.Println("The count is 12.")
  default :
    fmt.Println("The count is < 12.")  	

  }


}
