package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := []int{45, 50, 23, 10, 56, 32, 21, 5}
	sort.Ints(ages) //Unlike the strings methods this does changes the original array or slice.
	fmt.Println(ages)

	index := sort.SearchInts(ages, 21) //This returns the index of the searched value in the ages slice ie 21

	fmt.Println(ages[index], index)

	unfoundIndex := sort.SearchInts(ages, 32) //This returns: length of age slice + 1 if the searching parameter is larger than all the values or the value that is closest to the search parameter.
	fmt.Println(unfoundIndex)

	// Sort for String
	names := []string{"Romulus=Quirinus", "Orion", "Merlin", "Hassan of the Mountains", "Solomon", "Noah"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "Quirinus"))
}
