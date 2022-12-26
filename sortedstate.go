package main

import (
	"fmt"
	"sort"
)

func main() {
	state := make(map[string]string)
	state["Delhi"] = "Delhi"
	state["Kerala"] = "Thiruvananthapuram"
	state["Tamil Nadu"] = "Chennai"
	state["Maharashtra"] = "Mumbai"
	state["Uttar Pradesh"] = "Lucknow"

	stateSlice := make([]string, 0, len(state))
	for state1 := range state {
		stateSlice = append(stateSlice, state1)
	}

	sort.Strings(stateSlice)

	for _, state1 := range stateSlice {
		fmt.Printf("%s: %s\n", state1, state[state1])
	}
}
