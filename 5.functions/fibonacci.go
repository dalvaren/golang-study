package main

import "fmt"

func main() {
	fib := fibonacci(7)
	fmt.Println(fib)
}

func fibonacci(n int) (res int){
	if n <= 1 {
		res = 1
		return
	}
	res = fibonacci(n - 1) + fibonacci(n - 2)
	return
}