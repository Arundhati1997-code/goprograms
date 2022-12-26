package main

import "fmt"

type Friend struct {
	name            string
	mobNumber       string
	alternateMobNum string
	address         string
	city            string
	collegeFriend   bool
}

func main() {

	friend1 := Friend{"Priyanka", "98657890", "34768889", "Township", "Delhi", false}
	friend2 := Friend{"Rizwaan", "987878765", "", "hitech", "Hyderbaad", false}
	friend3 := Friend{"Sufiyan", "434456567", "34768889", "ojhar", "Nashik", true}
	friend4 := Friend{"Mac", "78664799", "34768889", "laxmiRoad", "Pune", true}
	friend5 := Friend{"Ken", "987656649", "", "College-Road", "Nagpur", true}
	friend6 := Friend{"Jenny", "999556789", "", "Airforce", "Pune", false}
	friend7 := Friend{"Diksha", "964336777", "34768889", "camp", "Mumbai", false}
	friend8 := Friend{"Tom", "789089995", "", "fish-Maeket", "Kokan", false}
	friend9 := Friend{"henrry", "078978900", "34768889", "camp2", "lonawalaa", false}
	friend10 := Friend{"Albert", "098080786", "34768889", "Hinjawadi", "Jaipur", true}

	Details := []Friend{
		friend1,
		friend2,
		friend3,
		friend4,
		friend5,
		friend6,
		friend7,
		friend8,
		friend9,
		friend10,
	}

	for _, friendDetails := range Details {
		if friendDetails.collegeFriend == true {
			fmt.Println(friendDetails.name, "|", friendDetails.city)
		}
	}
}
