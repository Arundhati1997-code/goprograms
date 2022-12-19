package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	min := 10
	max := 50
	var nums int
	var avg float64
	count := 0
	for i := 1; i <= 10; i++ {
		nums = rand.Intn(max-min+1) + min
		fmt.Println(nums)
		count++
		avg = float64(nums) / float64(count)
	}

	fmt.Println(avg)
}
