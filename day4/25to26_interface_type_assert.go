package main

import "fmt"


type Student struct {
	name string
	id int
}



func main() {


	container := make([]interface{},3)

	container[0] = 1
	container[1] = "Hello"
	container[2] = Student{"student02", 2}

    // interface type assert by if
	for index, data := range container {
		if value, ok := data.(int); ok == true {
			fmt.Println("type is int . index is ", index, value)
		}

		if value, ok := data.(string); ok == true {
			fmt.Println("type is string", index, value)
		}

		if value, ok := data.(Student); ok == true {
			fmt.Println("type is Student", index, value)
		}
	}


	// interface type assert by switch
    for index, data := range container {

    	switch value := data.(type) {
    	case int: 
    		fmt.Println("by switch ::: type is int . index is ", index, value)
        case string:
        	fmt.Println("by switch ::: type is string", index, value)
        case Student:
        	fmt.Println("by switch ::: type is Student", index, value)
        }

    }


}
