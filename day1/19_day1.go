
package main


import "fmt"


func main() {

	var ch byte
	ch = 97
	fmt.Println("ch=", ch)
	fmt.Printf("ch=%c\n", ch)
	fmt.Printf("ch=%d\n", ch)

	ch = 'c'
	fmt.Printf("ch=%d,ch=%c\n",ch,ch)

	// Capital letters to Lowercase letters 32
	// a:97 >> A:65
	ch1,ch2 := 'a','B'
	fmt.Printf("ch1=%c,ch2=%c\n",ch1-32, ch2+32)

    //The escape characters
    fmt.Printf("This is the GO %c", '\n')
    fmt.Printf("This is the JAVA.")
}