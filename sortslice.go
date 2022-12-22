package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Create an empty slice to store the integers
	var slice []int

	// Create a scanner to read input from the user
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter integers to add to the slice. Enter 'X' to quit.")

	// Loop until the user enters 'X'
	for {
		// Prompt the user to enter an integer
		fmt.Print("Enter an integer: ")
		scanner.Scan()
		input := scanner.Text()

		// Check if the us
		if strings.ToUpper(input) == "X" {
			break
		}

		// Convert the input to an integer
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer or 'X' to quit.")
			continue
		}

		// Add the integer to the slice
		slice = append(slice, num)

		// Sort the slice
		sort.Ints(slice)

		// Print the contents of the slice
		fmt.Println("Slice:", slice)
	}
}
