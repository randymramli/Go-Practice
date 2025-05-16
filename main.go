package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Create a new bill name: ", reader)
	b := newBill(name)
	fmt.Println("Created the bill -", b.name)

	return b
}

func prompOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip :)  ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("price must be a number")
			prompOptions(b)
		}
		b.addItem(name, p)
		fmt.Println("Item added -", name, price)
		prompOptions(b)

	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)
		p, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("tip must be a number")
			prompOptions(b)
		}
		b.addTip(p)
		fmt.Println("Tip added -", tip)
		prompOptions(b)

	case "s":
		fmt.Println("you chose to save the bill", b)
		a := b.format()
		fmt.Println(a)

	default:
		fmt.Println("Not a valid option...")
		prompOptions((b))
	}
}

func main() {
	mybill := createBill()
	prompOptions(mybill)
	// fmt.Println(mybill)
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
