package main

import (
	"fmt"
	"errors"
)

func main() {
	fmt.Println("Starting application...")
	fmt.Println()

	errNotFound := errors.New("Not found error")
	fmt.Printf("Error: %v\n", errNotFound)
	fmt.Println("---")

	if _, err := BiggerThanTwo(1); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println("---")
	if _, err := BiggerThanTwo(4); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println("---")

	if _, err := LowerThanTwo(4); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println("---")
}

func BiggerThanTwo(number int) (int, error) {
	if number < 3 {
		return number, errors.New("Number shall be bigger than 2")
	}
	return number, nil
}

func LowerThanTwo(number int) (int, error) {
	if number > 1 {
		return number, fmt.Errorf("The number %d shall be lower than 2", number)
	}
	return number, nil
}
