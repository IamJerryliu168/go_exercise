package main

import "fmt"
import "math/rand"
import "time"


func InitSlice(s []int) {
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<len(s);i++{
		s[i]=rand.Intn(100) //random number < 100
	}

}

func BubbleSOrt(s []int) {
    n := len(s)
	for i:=0;i<n-1;i++{
		for j:=0;j<n-i-1;j++{
			if s[j]>s[j+1]{
				s[j],s[j+1] = s[j+1],s[j]
			}
		}
	}
}

func main() {
	s := make([]int, 10)
	InitSlice(s)
	fmt.Println("before bubble sort, s = ", s)
	BubbleSOrt(s)
	fmt.Println("after bubble sort, s = ", s)
}
