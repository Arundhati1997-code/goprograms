package main

import "fmt"

func main() {

	var fact, num int
	fact = 1

	fmt.Print("Enter any Number to find the Factorial:")
	fmt.Scanln(&num)

	for i := 1; i <= num; i++ {
		fact *= i
	}
	fmt.Println("The Factorial of ", num, " = ", fact)
}
