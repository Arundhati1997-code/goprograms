package main

import "fmt"

func main() {
	numbers := [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	numb := numbers[4:14] //created sliceo of arr

	isOdd := func(x int) bool { //the use element of slice to check odd number
		if x%2 == 1 {
			return true
		}
		return false
	}

	isEven := func(x int) bool { //the use element of slice to check even number
		if x%2 == 0 {
			return true
		}
		return false
	}

	oddNumbers := []int{}
	for _, num := range numb { //then access odd values from isoddfunc to put in oddnumbersslice
		if isOdd(num) {
			oddNumbers = append(oddNumbers, num)
		}
	}

	evenNumbers := []int{}
	for _, num := range numb { //then access odd values from isoddfunc to put in evennumbersslice
		if isEven(num) {
			evenNumbers = append(evenNumbers, num)
		}
	}

	fmt.Println("The slice from array:", numb)
	fmt.Println("Odd numbers slice:", oddNumbers)
	fmt.Println("Even numbers slice:", evenNumbers)
}
