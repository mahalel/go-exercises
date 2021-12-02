package main

import (
	"fmt"
	"os"

	days "go/exercises/days"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please enter a day")
		return
	}
	whichDay(os.Args[1])
}

func whichDay(day string) {
	switch day {
	case "1":
		fmt.Println("Running day 1")
		days.One()
	case "2":
		fmt.Println("Running day 2")
		days.Two()

	default:
		fmt.Println("Pick a day between 1 and 25")
	}
}

//https://www.geeksforgeeks.org/go-operators/?ref=lbp
//https://learnxinyminutes.com/docs/go/
//https://www.golang-book.com/books/intro/6
//https://www.golangprograms.com
