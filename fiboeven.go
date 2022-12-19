package main

import (
	"fmt"
)

func Fibonacci(n int) (int, int) {

	n1 := 0
	n2 := 1
	count := 0
	var n3 int
	for i := 1; i <= n; i++ {
		fmt.Println(n1)
		n3 = n1 + n2

		n1 = n2

		n2 = n3
		if n3%2 == 0 {
			count++
		}
	}
	return n3, count

}

func main() {
	var n int
	fmt.Println("Enter the number")
	fmt.Scan(&n)
	n3, count := Fibonacci(n)
	fmt.Println(n3, count)

}
