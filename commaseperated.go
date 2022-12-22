package main

import (
	"fmt"
	"strings"
)

func main() {
	var numbers string
	fmt.Println("Enter the numbers:")
	fmt.Scanln(&numbers)
	num := strings.Split(numbers, ",")
	fmt.Println("The numbers are:", num)
	numb := num[0:len(num)]
	fmt.Println("The numbers are:", numb)
}
