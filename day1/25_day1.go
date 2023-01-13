package main

import "fmt"

func main() {
  
  var flag bool
  flag = bool(false)
  //flag = bool(1)  // cannot convert 1 (untyped int constant) to type bool
  fmt.Printf("flag is %T, value is %t\n", flag, flag)

  var ch byte
  ch = 'a'
  var t int
  t = int(ch)
  fmt.Printf("t is %d", t)    
}	