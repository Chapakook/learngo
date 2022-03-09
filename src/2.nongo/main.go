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

	// range arr -> [] [] ... result = index,value
	for _, work := range works {
		startWork(work)
	}
}

func startWork(work string) {
	fmt.Println(work, "Start!!")
	// Work Something ...
	time.Sleep(time.Second * 5)
	fmt.Println(work, "Done!!")
}
