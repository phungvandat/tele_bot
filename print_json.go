package main

import (
	"encoding/json"
	"fmt"
)

func PrintJSON(val interface{}) {
	b, _ := json.MarshalIndent(val, "", "\t")
	fmt.Println(string(b))
}
