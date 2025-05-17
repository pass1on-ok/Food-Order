package main

import (
	"fmt"
)

var subTotalBill float64

var customerOrder = make(map[string]uint, 0)

func main() {
	var customerName string
	fmt.Println("What is your first name?")
	fmt.Scan(&customerName)

	greet(customerName)
	orderItems()
	displayGeneratingBill()
	generateBill()
	modifyOrder()
	printFinalBill()
	sayTata(customerName)
}
