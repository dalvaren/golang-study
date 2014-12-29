package main

import "fmt"
import "strings"

func main() {
	var fullName string = "Daniel de Alvarenga Campos"

	fmt.Println(fullName)
	fmt.Println(strings.TrimSpace(fullName))
	fmt.Println()

	var names = strings.Fields(fullName)
	for _, name := range names {
		printOnlyIfNameIsBiggerThanTwoChars(name)
	}
	fmt.Println()

	names = strings.Split(fullName, " ") // Same thing than Fields, but specifying a sepparator
	for _, name := range names {
		printOnlyIfNameIsBiggerThanTwoChars(name)
	}
	fmt.Println(strings.Join(names, "-"))
	fmt.Println()

	var frase string = "PreCacheiro"
	fmt.Println(frase)
	fmt.Println(strings.HasPrefix(frase, "Pre"))
	fmt.Println(strings.HasSuffix(frase, "eiro"))
	fmt.Println(strings.Contains(frase, "cache"))
	fmt.Println(strings.Index(frase, "Cache"))
	fmt.Println()

	fmt.Println(strings.Replace(frase, "Pre", "Marmita", -1))
	fmt.Println(strings.Count(frase, "Pre"))
	fmt.Println(strings.ToLower(frase))
	fmt.Println(strings.ToUpper(frase))
}

func printOnlyIfNameIsBiggerThanTwoChars(name string) {
	if(len(name) > 2){
		fmt.Printf("%s\n", name)
	}
}