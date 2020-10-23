package main

import (
	models2 "apollo-adminserivce/internal/app/portal/models"
	"encoding/json"
	"fmt"
)

func main() {
	a := models2.Item{}
	j, _ := json.Marshal(a)
	fmt.Println(string(j))
}
