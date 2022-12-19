package main

import (
	"fmt"
	"math/rand"
	"time"
)

func gamingsimulation(stake int, goal int, trial int) {
	totalwin := 0
	totalbets := 0
	for i := 0; i < trial; i++ {
		betmade := 0
		currentstake := stake
		for currentstake > 0 && currentstake < goal {
			betmade++
			rand.Seed(time.Now().UnixNano())
			outcome := rand.Intn(2)
			if outcome == 1 {
				currentstake++
			} else {
				currentstake--
			}
		}
		if currentstake == goal {
			totalwin++
		}
		totalbets += betmade
	}

	winper := float64(totalwin) / float64(trial) * 100
	winavg := float64(totalwin) / float64(trial)

	fmt.Println("The number of wins:", totalwin)
	fmt.Println("The number of bets:", totalbets)
	fmt.Println("The number of win percentage:", winper)
	fmt.Println("The average win average:", winavg)

}
func main() {
	gamingsimulation(10, 20, 5)
}
