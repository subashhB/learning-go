package main

import "fmt"

var score = 99.5

func main() {

	// ! At this point the score would be undefined because it is not inside package score at this point but inside the block scope.
	// var score = 99.5

	// !You need to make sure that you run both the files in the terminal otherwise these imported values and functions would be undefined.
	sayHello("Mario")

	showScore()
	for _, v := range points {
		fmt.Println(v)
	}

}
