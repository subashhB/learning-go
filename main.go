package main

import "fmt"

func main() {
	// Print
	// This type of print doesn't add new line so the console print would be "hello, world"
	fmt.Print("hello, ")
	fmt.Print("world")
	fmt.Print("\nNew Line\n") // This is how you can print new line.

	// Println
	// This would print these two statements in new lines.
	fmt.Println("Hello")
	fmt.Println("Duniya")

	// Printing variables using print function
	age := 32
	name := "naam"

	fmt.Println("my name is ", age)
	fmt.Println("and my name is ", name)

	// Formatted string
	// * %v is the default format specifier
	// There are different format specifier
	fmt.Printf("My age is %v and my name is %v \n", age, name)
	fmt.Printf("My age is %q and my name is %q \n", age, name) // %q puts the formatter into quotation
	// ! This will not work for numbers and only string. Output: My age is ' '

	// %T will print the type of the variable.
	fmt.Printf("Age is of type %T\n", age)
	// * Output: Age is of type int

	// Printing float
	fmt.Printf("you scored %f points.\n", 22.3)
	fmt.Printf("you scored %.2f points.\n", 22.3)

	// Sprintf: Saves the formatted string
	var str = fmt.Sprintf("My age is %q and my name is %q \n", age, name)
	fmt.Printf("The saved String is: %v", str)

}
