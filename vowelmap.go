package main

import "fmt"

func countVowls(s string) (int, map[rune]int) {
	vowelDist := make(map[rune]int)

	numVowls := 0

	for _, r := range s {
		if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' {
			numVowls++

			vowelDist[r]++
		}
	}

	return numVowls, vowelDist
}
func main() {
	numVowls, vowelDist := countVowls("I am learning GoLang")

	fmt.Println("Number of vowels:", numVowls)
	fmt.Println("Vowel distribution:", vowelDist)

}
