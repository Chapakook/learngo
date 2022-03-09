package main

import (
	"fmt"
	"time"
)

func main() {
	works := []string{
		"1st work",
		"2nd work",
		"3rd work",
		"4th work",
		"5th work",
		"6th work",
		"7th work",
		"8th work",
		"9th work",
		"10th work",
	}

	// channal
	c := make(chan string)

	// range arr -> [] [] ... result = index,value
	for _, work := range works {
		go startWork(work, c)
	}

	// count go work
	for i := 0; i < len(works); i++ {
		fmt.Println(<-c)
	}
}

func startWork(work string, c chan<- string) {
	fmt.Println(work, "Start!!")
	// Work Something ...
	time.Sleep(time.Second * 5)
	c <- work + " Done!!"
}
