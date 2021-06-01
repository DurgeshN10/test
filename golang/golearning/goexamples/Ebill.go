package main

import "fmt"

func main() {

	var units, surCharge, amount, totAmount float64

	fmt.Print("Enter the Consumed Units = ")
	fmt.Scanln(&units)

	if units > 500 {
		amount = units * 12.65
		surCharge = 125
	} else if units >= 300 {
		amount = units * 10.75
		surCharge = 100
	} else if units >= 200 {
		amount = units * 8.26
		surCharge = 85
	} else if units >= 100 {
		amount = units * 5.98
		surCharge = 65
	} else {
		amount = units * 3.85
		surCharge = 45
	}
	totAmount = amount + surCharge
	fmt.Println("Electricity Bill = ", totAmount)
}
