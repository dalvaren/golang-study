package main

import "fmt"

func main() {
	map1 := make(map[string]int)
	map1["key1"] = 2
	map1["key2"] = 3
	fmt.Println(map1)
	fmt.Println()

	_, existValue := map1["key3"]
	fmt.Println("Exist value in map:", existValue)
	fmt.Println(map1)
	fmt.Println()

	map1["key4"] = 6
	map1["key5"] = 7
	fmt.Println(map1)
	delete(map1, "key4")
	fmt.Println(map1)
	fmt.Println()	
}