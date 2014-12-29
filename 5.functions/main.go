package main

import "fmt"

func main() {
	fmt.Println("Functions")
	fmt.Println()

	customFunc := testFunc
	customFunc()
	fmt.Println()

	var name string = "John"
	fmt.Println(name)
	changeName(&name)
	fmt.Println(name)
	fmt.Println()

	variadicFunc("Daniel", "Alvarenga", "Campos")
	fmt.Println()

	callerFunc()
	fmt.Println()

	func() {
		fmt.Println("Anonymous function.")
		}()
}

func testFunc() {
	fmt.Println("Test func called from reference")
}

func changeName(name *string) {
	*name = "Doe"
}

func variadicFunc(names ...string) {
	if len(names) > 0 {
		for _, name := range names {
			fmt.Println(name)
		}
	}	
}

func callerFunc() {
	fmt.Println("1")
	defer defferedFunc()
	fmt.Println("3")
}

func defferedFunc() {
	fmt.Println("2")
}