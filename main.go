package main

import (
	"fmt"
	"strings"
)

func main() {
	greetings := "Hi Mom"
	// Strings is the library that contains methods related to stringn
	fmt.Println(strings.Contains(strings.ToLower(greetings), "hi"))
	fmt.Println(strings.ReplaceAll(greetings, "Hi", "Hello"))
	fmt.Println(strings.ToUpper(greetings))
	fmt.Println(strings.Index(greetings, "om")) //returns the position where 'om' occurs, in this case 4
	fmt.Println(strings.Split(greetings, " "))  //returns the split strings into slice.

}
