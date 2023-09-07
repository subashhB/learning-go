package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, reader *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin) // os.Stdin means the standard input from the terminal

	// fmt.Print("Create a new bill name: ")
	// name, _ := reader.ReadString('\n') // inplace of _ there is supposed to be error but we don't need that at the moment.
	// // * '\n' in ReaderString means that to submit the reader input when the user presses 'Enter'
	// name = strings.TrimSpace(name)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose Options (a - add item, s - save, t - add tip): ", reader)
	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)
		fmt.Println(name, price)
	case "s":
		fmt.Println("Save Bill")
	case "t":
		tip, _ := getInput("Enter tip amount (Rs.): ", reader)
		fmt.Println(tip)
	default:
		fmt.Println("Invalid Input")
		promptOptions(b)
	}

}

func main() {
	myBill := createBill()
	promptOptions(myBill)
	fmt.Println(myBill)
}
