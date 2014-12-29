package main

import "fmt"

func main() {
	myName := "John"
	var myPointer *string

	myPointer = &myName

	fmt.Printf("value of myName at start: %s \n", myName)
	fmt.Printf("address of myPointer: %p \n", myPointer)
	fmt.Printf("value of myPointer: %s \n", *myPointer)
	fmt.Printf("Changing the pointer value... \n\n")

	*myPointer = "Teresa"

	fmt.Printf("value of myName: %s \n", myName)
	fmt.Printf("value of myPointer: %s \n", *myPointer)
}