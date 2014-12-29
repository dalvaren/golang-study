package main

import "fmt"
import "sort"

func main() {
	// var firstArray [5]int
	var firstArray = [5]int{1,2,3,4,5}
	// firstArray := new([5]int)
	firstArray[0] = 2
	fmt.Println(firstArray)
	fmt.Println()

	// var s []int = firstArray[2:5]
	// var s []int = firstArray[:]
	var s []int = firstArray[1:3]
	fmt.Println(s)
	fmt.Println()

	s = s[:4]
	fmt.Println(s)
	fmt.Println()

	sum := SumSlice(firstArray[2:5])
	fmt.Println(sum)
	fmt.Println()

	var s2 []int = make([]int,5,10)
	fmt.Println(s2)
	fmt.Println()

	s2 = append(s2, 3, 2, 1)
	fmt.Println(s2)
	fmt.Println()

	sort.Ints(s2)
	fmt.Println("s2 is sorted: ", sort.IntsAreSorted(s2))
	fmt.Println(s2)
	fmt.Println("s2 is sorted: ", sort.IntsAreSorted(s2))
	fmt.Println(sort.SearchInts(s2,2))
	fmt.Println()
}

func SumSlice(numbers []int) int{
	sum := 0
	for _, number := range numbers {
		sum = sum + number
	}
	return sum
}