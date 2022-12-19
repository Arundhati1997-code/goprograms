package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var toss int
	rand.Seed(time.Now().UnixNano())
	toss = rand.Intn(2)
	if toss == 0 {
		fmt.Println(toss)
		fmt.Println("Head")

	} else {
		fmt.Println(toss)
		fmt.Println("Tail")
	}

}
