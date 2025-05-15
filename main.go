package main

import "fmt"

func main() {

	myBill := newBill("Mario's Bill")
	myBill.addItem("Shrimp Salad", 14.99)
	myBill.addItem("Pie", 5.99)
	myBill.addItem("Coke Float", 4.99)
	myBill.addTip(3.50)
	fmt.Println(myBill)
	fmt.Println(myBill.format())

}

// func pointers() {
// 	name := "haha"

// 	m := &name

// 	fmt.Println("memory address: ", m)
// 	fmt.Println(name)

// 	updateName(m)
// 	fmt.Println(name)

// }

// func updateName(x *string) {
// 	*x = "wedge"
// }
