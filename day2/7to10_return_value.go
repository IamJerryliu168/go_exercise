package main

import "fmt"


func MyFunc01() string {
	return "MyFunc01()"
}

func MyFunc02() (str1, str2, str3 string) {
	str1, str2, str3 = "MyFunc02-str1","MyFunc02-str2","MyFunc02-str3"
	return
}

func MyFunc03()(a int, str string) {
	a, str = 8, "MyFunc03-str"
	return
}

func MaxMin(a int, b int) (min,max int) {
	if a > b {
		min = b 
        max = a
	} else {
		min = a
		max = b
	}
	return
}


func main() {
	f1 := MyFunc01()
	fmt.Printf("The f1 resut %s\n", f1)

	f2_1,f2_2,f2_3 := MyFunc02()
	fmt.Printf("The f2 result is f2_1=%s,f2_2=%s,f2_3=%s\n",f2_1, f2_2, f2_3)

	f3_1,f3_2 := MyFunc03()
    fmt.Printf("The f3 result is f3_1=%d,f3_2=%s\n",f3_1, f3_2)	

    min,max:=MaxMin(10,20)
    fmt.Printf("The min is %d, the max is %d\n",min, max)	

}