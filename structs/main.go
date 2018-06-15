package main

import "fmt"

type person struct {
	firstName string
	lastName string
}

func main() {
	// The risky way to define
	alex := person {"Alex", "Anderson"}

	// The more verbose way
	bob := person {firstName:"Bob", lastName:"Hendrix"}

	fmt.Println(alex, bob)

}