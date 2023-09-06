package main

import "fmt"

func updateName(name string) string {
	name = "wedge"
	return name //If we have to get this value then we return the value
}

func updateMenu(menuMap map[string]float64) {
	menuMap["coffe"] = 29.9
}

func main() {
	// Non pointer values
	name := "Tifa"

	updatedName := updateName(name) //Here a copy of the value is sent not the actual value so when something happens inside the function it happens to the copy and not the actual variable.

	fmt.Println(name)
	fmt.Println(updatedName)

	// Pointer Wrapper Values
	menu := map[string]float64{
		"soup":           44.9,
		"pie":            79.9,
		"salad":          67.7,
		"toffee pudding": 35.5,
	}

	updateMenu(menu) //In this case the original value does change for array, slice, maps, struct
	// This happens so because in these datatypes Go splits the data into two memory blocks.
	// One block contains the values and other contains "Pointer" pointing to the memory address to the values.
	// When passed by value it copies the memory address itself so whatever the changes are done are reflected to the original values since the copy points to the same memory location.
	fmt.Println(menu)

}
