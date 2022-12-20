package main

import (
	"fmt"
)

func main() {
	numbers := [...]int{34, 67, 55, 33, 12, 98, 88, 76}
	numb := numbers[0:8]
	fmt.Println(numb)

}
