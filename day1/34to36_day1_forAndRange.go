package main

import "fmt"
import "time"

func main() {

	str := "This is the go exercise!"

    // 1
	for i := 0; i < len(str); i++ {
		fmt.Println("The index is ", i, ".The value is", str[i])
	}

	// 2
	for i, value := range str {
		fmt.Printf("The index %d, the value %c\n", i, value)
	}

	// 3
	for _, value := range str {
		fmt.Println("The value is", value)
	}


	// break and continue
	j := 0
	for {
		j++
		time.Sleep(time.Second)
		if j > 5 {
			break
			//continue
		}
		fmt.Println("j = ", j)
	}
}
