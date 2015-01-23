package main

import (
	"fmt"
	"myPackage"
)

func main() {
	fmt.Println("Starting application...")
	if _, err := myPackage.Test(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println("...After test call")

}
