package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Starting application...")
	fmt.Println()

	params := ""
	if len(os.Args) > 1 {
		params += strings.Join(os.Args[1:],", ")
	}
	fmt.Printf("The params are: %s\n", params)
	fmt.Println("---")
}
