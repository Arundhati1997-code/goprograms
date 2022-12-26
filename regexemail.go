package main

import (
	"fmt"
	"regexp"
)

func main() {
	reg := regexp.MustCompile(`^[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z]{2,}$`)

	emails := []string{
		"user@example.com",
		"user123@example.com",
		"invalid@example",
		"invalid@.com",
	}

	for _, emid := range emails {
		fmt.Printf("%s: %v\n", emid, reg.MatchString(emid))
	}
}
