package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Println("enter the number:")
	fmt.Scan(&num)

	if num%3 == 0 && num%5 == 0 && num%7 == 0 {
		fmt.Println("PlingPlangPlong")
	} else if num%3 == 0 && num%5 == 0 {
		fmt.Println("PlingPlang")
	} else if num%3 == 0 && num%7 == 0 {
		fmt.Println("PlingPlong")
	} else if num%5 == 0 && num%7 == 0 {
		fmt.Println("PlangPlong")
	} else if num%3 == 0 {
		fmt.Println("Pling")
	} else if num%5 == 0 {
		fmt.Println("Plang")
	} else if num%7 == 0 {
		fmt.Println("Plong")
	} else {
		fmt.Println("Invalid")
	}
}
