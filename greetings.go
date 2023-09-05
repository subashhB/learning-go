// If we don't write package main here then we won't be able to share data and variable between the files
// Now after specifying this, go knows that both are the part of the same package
package main

import "fmt"

// Inside this file we don't have main function because there is only one main function ever.

var points = []int{20, 90, 35, 28, 47, 78}

func sayHello(name string) {
	fmt.Println("Hello ", name)
}

func showScore() {
	fmt.Println("Your score is ", score)
}
