package main

import "fmt"

func main() {
	var temp int
	fmt.Print("Enter the temperature in Celsius:")
	fmt.Scanf("%d", &temp)
	f := (float32(temp) * 1.8) + 32
	fmt.Println("Temperature in Fahrenheit is:", f)
}
