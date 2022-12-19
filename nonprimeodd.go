package main

import (
	"fmt"
)

func nonPrimeNumber(n int) bool {
	count2 := 0
	for j := 2; j < n; j++ {
		if n%j == 0 {
			count2++
		}
	}
	if count2 >= 1 {
		return true
	} else {
		return false
	}

}

func oddNonPrimeNumber(n int) int {
	if n%2 != 0 {
		return 1
	} else {
		return 0
	}
}

func main() {
	var n int
	count := 1
	fmt.Print("Enter any number: ")
	fmt.Scanln(&n)
	for i := 2; i < n; i++ {
		if nonPrimeNumber(i) == true {
			count += oddNonPrimeNumber(i)
		}
	}
	fmt.Println("Non Prime Numbers count:", count)
}
