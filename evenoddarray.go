package main

import "fmt"

func isOddcheck(numb []int) {
	oddNumbers := []int{}
	for i := 0; i < len(numb); i++ {
		if numb[i]%2 == 1 {
			oddNumbers = append(oddNumbers, numb[i])
		}

	}
	fmt.Println(oddNumbers)
}

func isEvenCheck(numb []int) {
	evenNumbers := []int{}
	for i := 0; i < len(numb); i++ {
		if numb[i]%2 == 0 {
			evenNumbers = append(evenNumbers, numb[i])
		}

	}
	fmt.Println(evenNumbers)
}

func main() {
	numbers := [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	numb := numbers[4:14] //created sliceo of arr
	isOddcheck(numb)
	isEvenCheck(numb)
}
