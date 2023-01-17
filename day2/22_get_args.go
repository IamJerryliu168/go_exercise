package main

import "fmt"
import "os"

func main() {

	params := os.Args

	fmt.Printf("the parameter counts is %d \n", len(params))

	for i:=0; i< len(params); i++ {
			fmt.Printf("for print:: the parameter parmeter[%d], value is %s \n", i, params[i])
	}

	for key,value := range params {
		fmt.Printf("range print:: the parameter parmeter[%d], value is %s \n", key, value)
	}

}	