package main

import "fmt"

func main() {
	go func() {
		fmt.Println("hello")
	}()
	fmt.Println("lihang")
	defer func() {
		fmt.Println("liji")
	}()
	test("a")("1")
}
func test(a string) func(string2 string) {
	return func(a string) {

	}
}
