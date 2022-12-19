package main

import "fmt"

func fact(num int) int {
	if num == 0 {
		return 1
	} else {
		return num * fact(num-1)
	}

}

func main() {

	var num int

	fmt.Print("Enter any Number to find the Factorial:")
	fmt.Scanln(&num)
	result := fact(num)
	fmt.Println("Factorial of ", num, " is ", result)

}
