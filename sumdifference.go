package main

import (
	"fmt"
)

func sum(n1 int, n2 int) int {
	return n1 + n2

}
func diff(n1 int, n2 int) int {
	return n1 - n2

}
func getsum_n_diff(n1 int, n2 int) (int, int) {
	sum := sum(n1, n2)
	diff := diff(n1, n2)

	return sum, diff

}
func main() {
	var n1, n2 int
	fmt.Println("Enter first number")
	fmt.Scan(&n1)
	fmt.Println("Enter second number")
	fmt.Scan(&n2)
	result := sum(n1, n2)
	fmt.Println("Sum of two numbers is ", result)
	result1 := diff(n1, n2)
	fmt.Println("Difference of two numbers is ", result1)
	sum, diff := getsum_n_diff(n1, n2)
	fmt.Println("Sum and diff of  two numbers is ", sum, diff)

}
