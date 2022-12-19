package main

import (
	"fmt"
)

func main() {
	var x, y, z int
	fmt.Println("Enter the side1 of the triangle")
	fmt.Scan(&x)

	fmt.Println("Enter the side2 of the triangle")
	fmt.Scan(&y)
	fmt.Println("Enter the side3 of the triangle")
	fmt.Scan(&z)

	if x == y && y == z && z == x {
		fmt.Print("The triangle is an equilateral triangle")
	} else if x == y && y != z || y == z && z != x || z == x && x != y {
		fmt.Println("The triangle is an isosceles triangle")
	} else {
		fmt.Println("The triangle is a scalene triangle")
	}
}
