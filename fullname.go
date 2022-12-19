package main

import (
	"fmt"
)

func main() {

	var firstname, lastname, fullname string

	fmt.Println("Enter first name: ")
	fmt.Scanln(&firstname)
	fmt.Println("Enter last name: ")
	fmt.Scanln(&lastname)
	fullname = firstname + " " + lastname
	fmt.Println("full name: ", fullname)
}
