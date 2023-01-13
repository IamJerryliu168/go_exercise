
package main

import "fmt"

func main() {

	var v1 complex64
	v1 = 3.2 + 10i
	fmt.Println("v1=", v1)
	v2 := 3.3 + 11i
    fmt.Println("v2=", v2)

    v3 := complex(3.4,12)
    fmt.Println("the real is",real(v3), "the virtual is", imag(v3))

}