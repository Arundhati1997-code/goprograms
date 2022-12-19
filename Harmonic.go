package main

import (
	"fmt"
)

func main() {
	var num int
	var sum float64 = 0
	fmt.Printf("Enter the number: ")
	fmt.Scanf("%d", &num)
	for i := 1; i <= num; i++ {
		sum += 1 / (float64)(num)

	}
	fmt.Println((sum))
}
