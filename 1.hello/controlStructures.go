package main

import "fmt"
import "runtime"

func main() {
	var test int = 2

	if test == 1 {
		fmt.Println("A")
	} else if test == 2 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	fmt.Println(runtime.GOOS)
	fmt.Println()

	fmt.Println(valueGreater(1,2))
	if value := valueGreater(4,3); value == 2 {
		fmt.Println("Result 1")
	} else {
		fmt.Println("Result 2")
	}
	fmt.Println()

	ret, err := checkDivisor(0)
	fmt.Println(ret, err)
	fmt.Println()	
}

func valueGreater(x int, y int) int {
	if x > y {
		return 1
	}
	return 2
}

func checkDivisor(x int) (number int, err bool) {
	if x == 0 {
		return 0, true
	}
	return x, false
}