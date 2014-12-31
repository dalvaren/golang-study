package main

import (
	"fmt"
	"player"
)

func main() {
	fmt.Println("Starting application...")
	fmt.Println()
	mainPlayer := player.NewPlayer("John", 100, 10)
	fmt.Println(mainPlayer)
}
