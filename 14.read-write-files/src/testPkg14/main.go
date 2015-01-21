package main

import (
	"fmt"
	"io/ioutil"
	"bufio"
	"os"
)

func main() {
	fmt.Println("Starting application...")
	fmt.Println()

	dataFile, err := ioutil.ReadFile("./test.dat")
	if err != nil {
		fmt.Println("Error reading the file.")
	}
	fmt.Print(string(dataFile))
	fmt.Println()
	fmt.Println("-----")

	fileHandler, err := os.Open("./test.dat")
	defer fileHandler.Close()
	if err != nil {
		fmt.Println("Error reading the file 2.")
	}
	bufferReader := bufio.NewReader(fileHandler)
	bufferData, err := bufferReader.Peek(5)
	fmt.Printf("5 bytes: %s\n", string(bufferData))
	fmt.Println("-----")

	f, err := os.Create("./dat2")
	if err != nil {
		fmt.Println("It was not possible to create the file!")
	}
	defer f.Close()
	f.WriteString("writes\n")
	fmt.Println("-----")

	outputFile, outErr := os.OpenFile("output.dat", os.O_WRONLY|os.O_CREATE, 0666)
	if outErr != nil {
		fmt.Println("It was not possible to create the file!")
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)
	outputWriter.WriteString("Buffered Text")
	outputWriter.Flush()

}
