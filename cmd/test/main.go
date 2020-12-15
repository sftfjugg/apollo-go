package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Ais struct {
		Name string
	}
	m := make(map[string]*Ais)
	a := new(Ais)
	a.Name = "123"
	m["l"] = a
	m["j"] = a
	s, _ := json.Marshal(m)
	fmt.Println(s)
	fmt.Println(string(s))
}
func test(a string) func(string2 string) {
	return func(a string) {

	}
}
