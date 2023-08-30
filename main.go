package main

import "fmt"

func main() {
	// strings 
	// Type Annotation
	var nameOne string = "Name"; //Implicitly typing a variable
	// ! We cannot use single quotes for the string, we must use double quotes.
	// !This does show error because we're declaring a variable but not using it.
	fmt.Println(nameOne);

	// Type Inference	
	var nameTwo = "Name Two";
	fmt.Println(nameTwo);

	// Just declaration
	var nameThree string;
	
	nameThree = "three"; // * If you don't have any value the console will just print blank
	fmt.Println(nameThree);

	// Short hand variable declaration
	nameFour := "Name Four"; // * This is equivalent to all the declaration above
	// ! You would use this only for the first time when you're declaring the variable but not when assigning value later on in the program.
	// ! You cannot use this outside of the function as well.
	fmt.Println(nameFour)

	// Numbers
	// * There are two type of number types: Integer and Floats
	// Integer
	var ageOne int = 10;
	var ageTwo  = 20;
	var ageThree int;
	ageThree = 30;
	ageFour := 40;

	fmt.Println(ageOne, ageTwo, ageThree, ageFour);

	// bits and memory
	var numOne int8 = 23; // Go allows us to set the range of ints like this in this case int8 ranges from -128 to 127
	// ! You cannot have 128 assigned in this case. var numOne int8 = 128 
	// There is concept of uint as well. Which doesn't allow negative number.

	var posOne uint = 2; // !Cannot allow -2 here
	fmt.Println(numOne, posOne)

	// float
	var scoreOne float32 = 1.5;
	var scoreTwo float64 = 1982739840701837410987908701928374901865897346509812394.128798341098374;
	scoreThree := 64.64; // * This will be inferred as default float64
	fmt.Println(scoreOne, scoreTwo, scoreThree) 

}