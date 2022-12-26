package main

import (
	"fmt"
	"regexp"
)

func insertCommaSpaces(s string) string {
	reg := regexp.MustCompile(",[^ ]")
	return reg.ReplaceAllString(s, ", ")
}
func main() {
	inp := "Q1 months are January,February,March."
	out := insertCommaSpaces(inp)
	fmt.Println(out)

}
