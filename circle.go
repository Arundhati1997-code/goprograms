package main

import "fmt"

func main() {
	var pi float64 = 3.142
	var dia int = 24
	rad := int(dia) / 2
	fmt.Println("Radius of circle is", rad)

	//area of circle
	area := pi * float64(rad) * float64(rad)
	fmt.Println(area)

	//area of circle
	perimeter := 2 * pi * float64(rad)
	fmt.Println(perimeter)
}
