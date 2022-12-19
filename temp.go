package main

import (
	"fmt"
)

func main() {
	var choice int
	var cel, fahr float64

	fmt.Println("Enter (0), it converts Celsius to Fahrenheit \nEnter (1),it converts Fahrenheit to celsius")

	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)
	switch choice {
	case 0:
		fmt.Print("Enter the value of Celsius: ")
		fmt.Scan(&cel)
		fahr = (cel * 9.0 / 5.0) + 32.0
		fmt.Println(fahr)
	case 1:
		fmt.Print("Enter the value of fahrenheit: ")
		fmt.Scan(&fahr)
		cel = (fahr - 32.0) * 5.0 / 9.0
		fmt.Println(cel)
	default:
		fmt.Println("Wrong choice")
	}

}
