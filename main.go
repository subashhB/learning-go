package main

import "fmt"

func main() {
	age := 45

	// Booleans
	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 45)
	fmt.Println(age != 50)

	if age <= 30 {
		fmt.Println("age is less than or equal to 30")
	} else if age <= 40 {
		fmt.Println("age is less than or equal to 40")
	} else {
		fmt.Println("age is less than or equal to neither 30 nor 40")
	}
	grandServants := []string{"Romulus=Quirinus", "Orion", "Merlin", "Hassan of the Mountains", "Solomon", "Noah"}

	for index, value := range grandServants {
		if index == 1 {
			fmt.Println("Continuing at pos ", index)
			continue //This will break the iteration but will continue the loop by taking the process to the start of the loop but will stop the process form going further down to the codes below.
			// @param 1 For index == 1, the print below iwll not be executed.
		}

		if index < 2 {
			fmt.Println("Break at position ", index)
			break //This breaks the loop entirely with no continue
		}

		fmt.Printf("The value at pos %v is %v\n", index, value)

	}

}
