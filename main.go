package main

import "fmt"

func main() {
	x := 0

	// While Loop
	for x < 5 {
		fmt.Println("Value of x is: ", x)
		x++
	}

	// For loop
	for i := 0; i < 5; i++ {
		fmt.Println("Value of i is: ", i)
	}

	// Loop through Slice
	names := []string{"Romulus=Quirinus", "Orion", "Merlin", "Solomon", "Hassan of the Mountains", "Noah"}
	for i := 0; i < len(names); i++ {
		fmt.Println("Name at", i, " of i is: ", names[i])
	}

	// For in loop
	for index, value := range names {
		fmt.Printf("The value at index %v is %v\n", index, value)
	}

	// * If you don't want to use the index, replace it with _
	//  * Same for the value
	for _, value := range names {
		fmt.Println("The name is ", value)
		value = "nothing" // *This won't change the value in the original slice because the value here is a local state.
	}

}
