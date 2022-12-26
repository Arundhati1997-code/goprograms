package main

import (
	"fmt"
)

func main() {
	var countryCapital = make(map[string]string)

	countryCapital["India"] = "New Delhi"
	countryCapital["Sri Lanka"] = "Colombo"
	countryCapital["Nepal"] = "Kathmandu"

	for country, capital := range countryCapital {
		fmt.Printf("Capital of %s is %s \n", country, capital)
	}
}
