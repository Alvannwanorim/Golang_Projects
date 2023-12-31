package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

func (b *bill) format() string {
	fs := "Billing breakdown for " + b.name + ": \n"
	var total float64 = 0

	for item, price := range b.items {
		fs += fmt.Sprintf("%-25v...$%v \n", item+": ", price)
		total += price
	}

	fs += fmt.Sprintf("%-25v...$%v \n", "tip: ", b.tip)

	fs += fmt.Sprintf("%-25v...$%0.2f \n", "Total: ", total+b.tip)

	return fs
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

func (b *bill) updateTip(tip float64) {
	b.tip = tip

}

func (b *bill) saveBill() {
	data := []byte(b.format())

	err := os.WriteFile(b.name+".txt", data, 0777)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bill saved ")
}
