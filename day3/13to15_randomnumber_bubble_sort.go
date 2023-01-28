package main

import "fmt"
import "math/rand"
import "time"

func main() {
	rand.Seed(time.Now().UnixNano())
	var a [5]int
	n := len(a)
	for i := 0; i < n; i++ {
		a[i] = rand.Intn(100) // rand number < 100
	}
	// print the array a before run bubble sort
	for key,value := range a {
		fmt.Printf("before order :: the key is %d, the value is %d \n", key, value)
	}
	// bubble sort
	for i:=0; i < n-1; i++ {
		for j:= 0; j < n-1-i; j++ {
			if a[j] > a[j+1] {
				a[j],a[j+1] = a[j+1],a[j]
			}
		}
	}
	for key,value := range a {
		fmt.Printf("after order :: the key is %d, the value is %d \n", key, value)
	}


}
