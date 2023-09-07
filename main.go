package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

		p, err := strconv.ParseFloat(price, 64) //64 is bit here, float64. err will capture the error if any occured like if we passed a string that is not a number string.
		if err != nil {
			fmt.Println("Price must be a number")
			promptOptions(b)
		}
		b.addItem(name, p)
		fmt.Println("Item added - ", name, price)
		promptOptions(b)
	case "s":
		fmt.Println("You chose to save the bill.")
		fmt.Println(b.format())
	case "t":
		tip, _ := getInput("Enter tip amount (Rs.): ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("Tip must be a number.")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("Tip added: ", tip)
		promptOptions(b)
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
