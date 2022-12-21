package main

import (
	"fmt"
	"strconv"
	"strings"
)

func max_secmax(numb []string) (float64, float64) {
	max := 0.0
	secmax := 0.0

	for i := 0; i < len(numb); i++ {
		num, _ := strconv.ParseFloat(numb[i], 64)
		if num > max {
			max = num

		}
	}

	for i := 0; i < len(numb); i++ {
		num, _ := strconv.ParseFloat(numb[i], 64)
		if num > secmax && num < max {
			secmax = num

		}
	}

	return max, secmax
}

func min_secmin(numb []string) (float64, float64) {

	min := 333333333.000
	secmin := 222222222.000

	for i := 0; i < len(numb); i++ {
		num, _ := strconv.ParseFloat(numb[i], 64)
		if num < min {
			min = num

		}
	}

	for i := 0; i < len(numb); i++ {
		num, _ := strconv.ParseFloat(numb[i], 64)
		if num < secmin && num > min {
			secmin = num

		}
	}

	return min, secmin
}

func main() {
	var numbers string
	fmt.Println("Enter the numbers:")
	fmt.Scanln(&numbers)
	num := strings.Split(numbers, ",")
	numb := num[0:len(num)]
	max, secmax := max_secmax(numb)
	min, secmin := min_secmin(numb)
	fmt.Println("The maximum number in the array is: ", max, secmax)
	fmt.Println("The minimum number in the array is: ", min, secmin)
}
