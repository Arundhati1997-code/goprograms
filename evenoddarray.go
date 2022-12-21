package main

import "fmt"

func main() {
	arr := [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	numb := arr[4:14]

	isOdd := func(x int) bool {
		if x%2 == 1 {
			return true
		}
		return false
	}

	isEven := func(x int) bool {
		if x%2 == 0 {
			return true
		}
		return false
	}

	oddNumbers := []int{}
	for _, num := range numb {
		if isOdd(num) {
			oddNumbers = append(oddNumbers, num)
		}
	}

	evenNumbers := []int{}
	for _, num := range numb {
		if isEven(num) {
			evenNumbers = append(evenNumbers, num)
		}
	}

	fmt.Println("The slice from array:", numb)
	fmt.Println("Odd numbers slice:", oddNumbers)
	fmt.Println("Even numbers slice:", evenNumbers)
}
