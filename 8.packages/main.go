package main

import (
	"fmt"
	"./test"
)
import _ "./test2"

func main() {
	fmt.Println("Starting app...")
	test.PrintTest()
}
