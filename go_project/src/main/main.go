package main

import (
	"fish"
	"fmt"
)

func main() {
	myArray = make([]int, 5, 10)

	fmt.Print("main package run\n")
	fish.TestPrint()
}
