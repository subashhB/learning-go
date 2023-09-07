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
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// Receiver Function: format the bill
// Now we can access this function by billObject.format() but we cannot call format() solo.
// * Another pros of using pointers is that we don't make copies everytime we call the function. If we call the method multiple times then it would make the copy that many times as well. Pointers save us from that.
func (b *bill) format() string {
	formattedString := "Bill Breakdown: \n\n"

	var total float64 = 0

	// List items
	for k, v := range b.items {
		formattedString += fmt.Sprintf("%-25v ....Rs.%v \n", k+":", v) // -25 here makes the string 25 characters long, even if the value isn't 25 characters long, white blank spaces are added.
		total += v
	}

	// Tips
	formattedString += fmt.Sprintf("\n%-25v ....Rs.%v", "total:", b.tip)

	// Total formatting in string
	formattedString += fmt.Sprintf("\n%-25v ....Rs.%.2f", "total:", total+b.tip) // IF we would have put +25 here then the spacing would have been on the left side of the character.

	return formattedString

}

// *Rule of thumb: Whenever we're calling a method to update the value we pass the pointer.

// Update the tip
func (b *bill) updateTip(tip float64) {
	// While working with the struct we don't need to dereference the pointer to get the value.
	b.tip = tip
}

// add an item to the bill
func (b *bill) addItem(itemName string, price float64) {
	b.items[itemName] = price
}
