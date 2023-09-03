package main

import (
	"fmt"
	"math"
)

func sayGreeting(name string) {
	fmt.Println("Good Morning", name)
}

func sayBye(name string) {
	fmt.Println("Good Bye", name)
}

func cycleNames(names []string, function func(string)) {
	for _, v := range names {
		function(v)
	}
}

func cirleArea(radius float64) float64 {
	return math.Pi * radius * radius
}

func main() {
	sayGreeting("Zelda")
	sayGreeting("Tifa")
	sayBye("Zelda")

	cycleNames([]string{"Romulus", "Orion", "Merlin", "Solomon", "Hasan", "Noah"}, sayGreeting) //We aren't calling the function otherwise we would be invoking the function here only which we don't want instead we are passing the reference of the function.
	cycleNames([]string{"Romulus", "Orion", "Merlin", "Solomon", "Hasan", "Noah"}, sayBye)

	area1 := cirleArea(7)
	area51 := cirleArea(7.5)
	fmt.Printf("Circle 1 is %.3f and Circle 2 is %.3f", area1, area51)

}
