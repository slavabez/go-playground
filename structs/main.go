package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName: "Party",
		contactInfo: contactInfo{
			zipCode: 12345,
			email: "jim@party.nex",
		},
	}

	fmt.Printf("%+v", jim)
}