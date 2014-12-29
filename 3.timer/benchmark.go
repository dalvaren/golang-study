package main

import "fmt"
import "time"

func main() {
	startBenchmarkTime := time.Now()
	defer func(){
		endBenchmarkTime := time.Now()
		fmt.Println("Total execution time: ", endBenchmarkTime.Sub(startBenchmarkTime))
		}()

	// fmt.Println("test of execution...")
}