package main

import "fmt"

func arrayaway(place [7]string) [7]int {
	var length [7]int
	for i := range place {
		length[i] = len(place[i])
		i++
	}
	fmt.Println("\n")
	return length
}

func main() {

	place := [...]string{"New York", "Paris", "London", "Sydney", "Tokyo", "Moscow", "Toronto"}
	result := arrayaway(place)
	fmt.Println(result)

}
