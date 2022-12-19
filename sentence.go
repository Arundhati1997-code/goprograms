package main

import (
	"bufio"
	"fmt"
	"os"
)

func vowelsAndConsonants(sent string) {
	vowCount := 0
	sp := 0
	for _, ch := range sent {
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
			vowCount++
		} else if ch == ' ' {
			sp++
		}
	}
	consCount := len(sent) - vowCount - sp
	fmt.Println("Vowels: ", vowCount)
	fmt.Println("Consonants: ", consCount)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a sentence: ")
	scanner.Scan()
	sent := scanner.Text()
	vowelsAndConsonants(sent)
}
