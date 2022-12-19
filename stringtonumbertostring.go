package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n string
	fmt.Print("Enter a number: ")
	fmt.Scanln(&n)

	res, _ := strconv.Atoi(n)
	res = res * 10
	fmt.Println("10 times the number:", strconv.Itoa(res))
}
