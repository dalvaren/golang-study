package main

import (
	"fmt"
	"time"
	"runtime"
)

func main() {
	fmt.Println("Starting application...")
	fmt.Println()

	runtime.GOMAXPROCS(2)

//	test1()
//	test2()
//	test3()
//	test4()
//	test5()
	test6()
	fmt.Println()
}

func test1() {
	go longWait()
	go shortWait()
	time.Sleep(10 * 1e9)
}

func longWait() {
	fmt.Println("Beginning long wait.")
	time.Sleep(5 * 1e9)
	fmt.Println("End long wait.")
}

func shortWait() {
	fmt.Println("Beginning short wait.")
	time.Sleep(2 * 1e9)
	fmt.Println("End short wait.")
}

func test2() {
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)
	time.Sleep(1 * 1e9)
}

func sendData(ch chan string) {
	ch <- "São Paulo"
	ch <- "Tokio"
	ch <- "New York"
	ch <- "Paris"
	ch <- "Hong Kong"
}

func getData(ch chan string) {
	for {
		fmt.Printf("%s ", <-ch)
	}
}

func test3() {
	go suck(pump())
	time.Sleep(1 * 1e9)
}

func pump() chan int {
	ch := make(chan int)
	go func() {
		for i:=0;;i++ {
			ch <- i
		}
	}()
	return ch
}

func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

func test4() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go pump1(ch1)
	go pump2(ch2)
	go suck2(ch1, ch2)
	time.Sleep(1 * 1e9)
}

func pump1(ch chan int) {
	for i:=0;;i++ {
		ch <- i*2
	}
}

func pump2(ch chan int) {
	for i:=0;;i++ {
		ch <- i+5
	}
}

func suck2(ch1 chan int, ch2 chan int){
	for i:=0;;i++{
		select {
		case value := <-ch1:
			fmt.Printf("%d. Received of 1: %d\n", i, value)
		case value := <-ch2:
			fmt.Printf("%d. Received of 2: %d\n", i, value)
		}
	}
}

func test5() {
	pending, done := make(chan int), make(chan int)

	go SendWork(pending)
	for i:=0; i < 5; i++ {
		go Worker(pending, done)
	}

	time.Sleep(1 * 1e9)
}

func Worker(in, out chan int) {
	for{
		t := <-in
		fmt.Printf("Processing: %d\n", t)
		t = t + 1
		out <- t
	}
}

func SendWork(ch chan int){
	for i:=0;;i++ {
		ch <- i
	}
}

func test4() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	go pump1(ch1)
	go pump2(ch2)
	go suck2(ch1, ch2)
	time.Sleep(1 * 1e9)
}
