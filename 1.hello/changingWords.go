package main

import "fmt"
import "strings"

func main() {
	fullName := "Daniel de Alvarenga Campos"
	names := strings.Fields(fullName)

	for index, name := range names {
		if len(name) < 3 {
			continue;
		}
		names[index] = strings.ToLower(name)
		fmt.Println(name)
	}

	fmt.Println()

	fmt.Println(strings.Join(names, "-"))
}