package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonBuf := `
	{
		"company": "A Co",
		"subjects": [
			"java",
			"c++",
			"go"
		],
		"isOk": true,
		"price": 666.6
	}
	`
	m := make(map[string]interface{}, 4)

	err := json.Unmarshal([]byte(jsonBuf), &m)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("tmp=%+v\n", m)

	//co := m["company"]
	//fmt.Println("company=", co)

	// below code not work
	//var str string
	//str = string(m["company"])
	//fmt.Println("str=", str)

	var str string
	for key, value := range m {
		//fmt.Printf("Key is %s,value is %s\n", key, value)
		switch data := value.(type) {
		case string:
			str = data
			fmt.Printf("map[%s] type is string and value is %s \n", key, str)
		case bool:
			fmt.Printf("map[%s] type is bool and value is %v \n", key, data)
		case float64:
			fmt.Printf("map[%s] type is float64 and value is %f \n", key, data)
		case []string:
			fmt.Printf("map[%s] type is []string and value is %v \n", key, data)
		case []interface{}:
			fmt.Printf("map[%s] type is []interface{} and value is %v \n", key, data)
		}
	}
}
