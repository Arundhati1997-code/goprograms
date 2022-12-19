package main

import (
	"fmt"
)

func main() {
	var principal, n, t, finalamt float64
	RoI1 := 6.25
	RoI2 := 7.50
	RoI3 := 9.5
	var r = RoI1 + RoI2 + RoI3

	fmt.Println("Enter the principal amonut:")
	fmt.Scan(&principal)

	fmt.Println("Enter the value of n:")
	fmt.Scan(&n)

	fmt.Println("Enter the time elapsed")
	fmt.Scan(&t)

	finalamt = principal * (1 + r/n) * n * t

	fmt.Println(finalamt)

}
