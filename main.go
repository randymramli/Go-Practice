package main

import "fmt"

func main() {

	myBill := newBill("Mario's Bill")

	fmt.Println(myBill)

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
