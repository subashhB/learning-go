package main

import "fmt"

// Now we're accepting pointer to the string as an arguement
func updateName(name *string) {
	// Take the pointer, derefernce it to see where it is pointing at and update the value at that memory block
	*name = "wedge"
}

func main() {
	// Non pointer values
	name := "Tifa"

	fmt.Println("Memory Address of the name: ", &name) //Here, &name is the pointer of the variable

	m := &name

	fmt.Println("Memory Address: ", m)              //Here m is the pointer
	fmt.Println("Value at the memory address:", *m) // *m allows us to access the value in the address that m is pointing to
	// The reference value for the value that pointer points to.

	fmt.Println(name)
	// To overcome the boundaries of Pass by value for the non pointer wrapping variables, we can use the pointers.
	// By the use of the pointers we can now change the value of the name variable, we weren't able to do so before.
	updateName(m)
	fmt.Println(name)
}
