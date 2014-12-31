// Package myLibrary is a test for a loaded library
package myLibrary

import (
	"fmt"
)

// PrintTest prints a test string in console
func PrintTest() {
	fmt.Println("Package myLibrary loaded!")
}

func SumNumbers(a, b int) int{
	return a + b
}
