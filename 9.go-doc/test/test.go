// Test package to perform tests in application
package test

import (
	"fmt"
)

// Initializes Test Package only printing "Initiating the test package"
func init() {
	fmt.Println("Initiating the test package")
}

// Print a test message
func PrintTest() {
	fmt.Println("This is a test")
}
