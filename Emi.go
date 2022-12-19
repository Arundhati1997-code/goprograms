package main

import (
	"fmt"
)

func main() {
	var principal, year, RoI, payment int
	var n = 12 * year

	fmt.Print("Enter the principal amount:")
	fmt.Scan(&principal)

	fmt.Print("Enter the number of year:")
	fmt.Scan(&year)

	fmt.Print("Enter the rate of interest:")
	fmt.Scan(&RoI)

	payment = principal*RoI/1 - (1 + RoI) ^ (-n)
	fmt.Print("The payment:", payment)

}
