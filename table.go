package main

import (
	"fmt"
)

func main() {
	var num, n, res int

	fmt.Print("Type a number: ")
	fmt.Scan(&num)
	fmt.Print("Type value of n: ")
	fmt.Scan(&n)
	if num >= 2 && num <= 25 {
		for i := 1; i <= n; i++ {
			res = num * i
			fmt.Println("num X i =", res)

		}
	}

}
