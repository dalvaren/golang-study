package main

import "fmt"
import "strconv"

func main() {
	var stringNumber string = "666"
	var myNumber int = 54

	fmt.Println(stringNumber)
	fmt.Println(myNumber)
	fmt.Println(strconv.IntSize)

	sum1, _ := strconv.Atoi(stringNumber)
	sum1 = sum1 + myNumber

	sum2 := strconv.Itoa(myNumber)
	sum2 = sum2 + stringNumber
	fmt.Println(sum2)
}