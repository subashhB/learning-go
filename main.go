package main

import (
	"fmt"
	"strings"
)

func getInitials(name string) (string, string) {
	s := strings.ToUpper(name)
	splits := strings.Split(s, " ")

	var initials []string

	for _, v := range splits {
		initials = append(initials, v[:1])
	}

	if len(initials) == 2 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}

func main() {
	firstInitial, secondInitial := getInitials("Tifa Lockhart")
	fmt.Println(firstInitial, secondInitial)

	firstInitial1, secondInitial1 := getInitials("Tifa Lockhart")
	fmt.Println(firstInitial1, secondInitial1)

	firstInitial2, secondInitial2 := getInitials("Barret")
	fmt.Println(firstInitial2, secondInitial2)
}
