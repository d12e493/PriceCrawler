package main

import (
	"fmt"
)

var test = "ABC"

func main() {
	fmt.Println(test)
	change()
	fmt.Println(test)
	change2()
}

func change() {
	test := "123"
	fmt.Println(test)
}

func change2() {
	test := "999"
	fmt.Println(test)
}
