package main

import "fmt"
import "time"

func main() {
	var t = time.Now()
	fmt.Println(t)
	fmt.Println(t.UTC())

	var week = 60 * 60 * 24 * 7 * 1e9// must be in nanosecs
	var weekFromNow = t.Add(time.Duration(week))

	fmt.Println(weekFromNow)

	fmt.Println(t.Format(time.RFC822))
	fmt.Println(t.Format(time.ANSIC))

	fmt.Println(t.Format("02 Jan 2006 15:04")) // pass an example
	fmt.Println(t.Format("20060102"))

	fmt.Printf("%02d-%02d-%04d %02d:%02d:%02d %d \n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond())
}