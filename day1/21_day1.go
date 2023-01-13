package main

import "fmt"

func main() {

	var ch byte
	var str string


    // the character only one character and using single quote marks
	ch = 'a'
    fmt.Println("ch=",ch)

	// 1: double quote marks
	// 2: one character or multi characters
	// 3: a '\0' char hidden in string
	str = "a"   // includer 'a' and '\0' 
    fmt.Println("str=",str)

    str = "This is the GO!"
    fmt.Printf("str[0]=%c,str[1]=%c\n",str[0],str[1])
}