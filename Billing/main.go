package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInputs(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	name, err := r.ReadString('\n')
	name = strings.TrimSpace(name)

	return name, err
}

func createNewBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := readInputs("Enter new bill name: ", reader)
	bill := newBill(name)

	return bill

}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := readInputs("Choose option: (a - add an item, s - save bill, t - add tip) ", reader)

	switch opt {
	case "a":
		name, _ := readInputs("Enter Item name: ", reader)
		price, _ := readInputs("enter Item price: ", reader)
		fmt.Println(name, price)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("Price must be a string")
			promptOptions(b)
		}
		b.addItem(name, p)
		promptOptions(b)

	case "s":
		b.saveBill()
		fmt.Println("Bill saved successfully", b)

	case "t":
		tip, _ := readInputs("Enter tip amount: ", reader)
		p, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("tip must be a string")
			promptOptions(b)
		}
		b.updateTip(p)
		promptOptions(b)

	default:
		promptOptions(b)
	}
}

func main() {
	bill := createNewBill()
	promptOptions(bill)

	fmt.Println(bill.format())

}
