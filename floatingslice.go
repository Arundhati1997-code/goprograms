package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func ComputeAverage(number []float64) float64 {
	sum := 0.0

	for _, num := range number {
		sum += num
	}

	return sum / float64(len(number))
}

func ComputeStandardDeviation(number []float64) float64 {
	sumOfSquares := 0.0

	average := ComputeAverage(number)

	for _, num := range number {
		sumOfSquares += (num - average) * (num - average)
	}

	return math.Sqrt(sumOfSquares / float64(len(number)))
}

func main() {
	var number []float64

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter a list of floating point numbers (enter 'q' or 'Q' to stop):")
	for scanner.Scan() {
		input := scanner.Text()

		if strings.ToLower(input) == "q" {
			break
		}

		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid floating point number.")
			continue
		}

		number = append(number, num)
	}

	average := ComputeAverage(number)
	standardDeviation := ComputeStandardDeviation(number)
	fmt.Printf("Average: %f\n", average)
	fmt.Printf("Standard deviation: %f\n", standardDeviation)
}
