package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := make(map[string]interface{})
	m["Company"] = "B.co"
	m["Subjects"] = []string{"java", "c++", "go"}
	m["IsOk"] = true
	m["Price"] = 666.6
	buf, err := json.Marshal(m)
	//buf, err:=json.MarshalIndent(m,""," ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf))
}
