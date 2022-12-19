package main

import (
	"fmt"
)

func Fibonacci(n int) int {

	if n <= 1 {

		return n

	}

	return Fibonacci(n-1) + Fibonacci(n-2)

}

func main() {
	Fibonacci(10)
	for i := 0; i < 10; i = i + 1 {

		fmt.Println(Fibonacci(i))

	}

}
