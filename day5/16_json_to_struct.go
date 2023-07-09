package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Comapny  string   `json:"company"`
	Subjects []string `json:"subjects"`
	IsOk     bool     `json:"isOk`
	Price    float64  `json:"price"`
}

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
	var tmp IT
	err := json.Unmarshal([]byte(jsonBuf), &tmp)

	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("tmp=", tmp)
	fmt.Printf("tmp=%+v\n", tmp)

	type IT2 struct {
		Subjects []string `json:"subjects"`
	}
	var tmp2 IT2
	err2 := json.Unmarshal([]byte(jsonBuf), &tmp2)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Printf("tmp=%+v\n", tmp2)

}
