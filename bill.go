package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
	total float64
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
		total: 0,
	}
	return b
}

func (b bill) format() string {
	fs := "Bill Breakdown: \n"
	// var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		b.total += v
	}

	b.total += b.tip

	fs += fmt.Sprintf("%-25v ...$%v\n", "Tip:", b.tip)

	fs += fmt.Sprintf("%-25v ...$%.2f", "Total:", b.total)

	return fs
}

func (b *bill) addTip(tip float64) {
	b.tip = tip
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

func (b *bill) save() {
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}
}
