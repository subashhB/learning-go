package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bill
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 59.9, "cake": 139.6},
		tip:   0,
	}

	return b
}

// Receiver Function: format the bill
// Now we can access this function by billObject.format() but we cannot call format() solo.
func (b bill) format() string {
	formattedString := "Bill Breakdown: \n"

	var total float64 = 0

	// List items
	for k, v := range b.items {
		formattedString += fmt.Sprintf("%-25v ....Rs.%v \n", k+":", v) // -25 here makes the string 25 characters long, even if the value isn't 25 characters long, white blank spaces are added.
		total += v
	}

	// Total formatting in string
	formattedString += fmt.Sprintf("\n%-25v ....Rs.%.2f", "total:", total) // IF we would have put +25 here then the spacing would have been on the left side of the character.

	return formattedString

}
