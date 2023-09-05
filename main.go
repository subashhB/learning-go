package main

import "fmt"

func main() {
	// menu is a key value pair.
	// menu := map[key_datatype]value_datatype
	menu := map[string]float64{
		"soup":           44.9,
		"pie":            79.9,
		"salad":          67.7,
		"toffee pudding": 35.5,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// Looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// Int as key type
	phonebook := map[int]string{
		9841984198:  "Shyam",
		98039830398: "Ram",
		9843984398:  "Sita",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[9841984198])

	// updating maps
	phonebook[9843984398] = "Alice"
	fmt.Println(phonebook)

	phonebook[98039830398] = "Bob"
	fmt.Println(phonebook)

	// This add a new data into maps
	phonebook[12345678] = "Alex"
	fmt.Println(phonebook)

}
