package main

import "fmt"

func main() {
	// * Arrays
	// Fixed length and cannot change that
	var age [3]int
	age = [3]int{2, 3, 4} // This is how you declare an array in Go.

	// Other assignment
	var ageTwo = [3]int{20, 35, 30}

	fmt.Println(age, ageTwo)
	fmt.Println("The length of the array age is ", len(age))

	// Short hand
	names := [4]string{"Romulus", "Orion", "Hassan", "Solomon"}
	names[3] = "Merlin"
	fmt.Println(names, len(names))

	// * Slices

	var scores = []int{100, 20, 30} //This is a slice if the length is not given

	// Unlike in arrays we can append a value to the slices.
	var newScores = append(scores, 85) // *But the append method doesn't append the value to the original array so it should be stored in new variable.
	scores = append(scores, 65)        //Or like this

	fmt.Println(newScores, scores)

	// * Ranges
	// This is a way to cut a range of elements form a array or slice [Somthing like slice in JavaScript]

	rangeOne := names[1:3]
	rangeTwo := names[1:]   //Go from index one to end of the array.
	rangeThree := names[:3] //From the first to the index of three but not including position 3.

	// since these slices are just ranges itself, we can easily append these as well.

	fmt.Println(rangeOne, rangeTwo, rangeThree)
	rangeOne = append(rangeOne, "Noah") //But this does effect the underlying arrays which is rangeTwo
	fmt.Println(rangeOne, rangeTwo, rangeThree)
}
