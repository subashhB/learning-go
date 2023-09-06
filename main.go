package main

import "fmt"

func main() {
	myBill := newBill("Luffy's Bill")
	fmt.Println(myBill.format())
}
