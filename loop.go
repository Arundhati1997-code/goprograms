package main

import "fmt"

func main() {
	var n int

	fmt.Println("Enter any number: ")
	fmt.Scanln(&n)

	for i := 0; i < 10; i++ {
		prod := n * i
		fmt.Println(prod)
	}
}
