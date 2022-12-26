package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	stateCapitals := map[string]string{
		"Jammu and Kashmir": "Srinagar",
		"Andhra Pradesh":    "Amaravati",
		"Arunachal Pradesh": "Itanagar",
		"Assam":             "Dispur",
		"Bihar":             "Patna",
		"Chhattisgarh":      "Raipur",
		"Goa":               "Panaji",
		"Gujarat":           "Gandhinagar",
		"Haryana":           "Chandigarh",
		"Himachal Pradesh":  "Shimla",
	}
	jsonInp, err := json.MarshalIndent(stateCapitals, "", " ")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(jsonInp)
	}

}
