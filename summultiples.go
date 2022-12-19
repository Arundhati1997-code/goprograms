package main

import "fmt"

func sumMultiples(num int) int {
	sum := 0
	for i := 1; i < num; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}
func main() {
	result := sumMultiples(20)
	fmt.Println(result)

}
