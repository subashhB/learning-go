package main

import "fmt"

func main() {
	myBill := newBill("Luffy's Bill")
	myBill.addItem("Onion Soup", 45.0)
	myBill.updateTip(50)
	fmt.Println(myBill.format())
}
