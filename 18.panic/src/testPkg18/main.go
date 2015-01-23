package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting application...")
//	panic("THIS IS PANIC!")
	test()
	fmt.Println("Continuing the program...")
	fmt.Println()
}

func test() {
	defer func(){
		if err := recover(); err != nil {
			fmt.Printf("Panicking: %s\n", err)
		}
	}()
	badCall()
	fmt.Println("After bad call...")
}

func badCall() {
	panic("bad end")
}
