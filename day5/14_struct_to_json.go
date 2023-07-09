package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Comapny  string
	Subjects []string
	IsOk     bool
	Price    float64
}

type IT2 struct {
	Comapny  string   `json:"company"`
	Subjects []string `json:"subjects"`
	IsOk     bool     `json:",string"`
	Price    float64  `json:",string"`
}

func main() {
	//object := IT{"A Co", []string{"java", "c++", "go"}, true, 666.6}
	object := IT2{"A Co", []string{"java", "c++", "go"}, true, 666.6}
	//buf,err := json.Marshal()
	buf, err := json.MarshalIndent(object, "", " ")

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf))
}
