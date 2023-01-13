
package main

import "fmt"

func main() {

    var ch byte
    var i int
    var str string
    fmt.Println("Pls enter ch's value:")
    fmt.Scanf("%c", &ch)
    fmt.Println("Pls enter str's value:")
    //fmt.Scanf("%s", &str)
    fmt.Scan(&str)
    fmt.Println("Pls enter i's value:")
    fmt.Scan(&i)
    //fmt.Scanf("%d", &i)

    fmt.Printf("ch is %T, i is %T, str is %T", ch, i, str)

    //fmt.scan(&ch)

}	