package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	// 1st
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)

	// 2nd
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// fmt.Println(alex)
	// fmt.Printf("%+v\n", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// jim.updateName("jimmy")
	// jim.print()

	// pointer explicit way
	// jimPointer := &jim
	// jimPointer.updateName("jimmy")
	// jim.print()

	// pointer shortcut way
	jim.updateName("jimmy")
	jim.contactInfo.email = "jim@outlook.com"
	jim.print()

	// slice
	greeting := []string{"Hi", "There", "How", "Are", "You"}
	updateGreeting(greeting) // without pointer operator like & *
	fmt.Printf("%+v\n", greeting)

	// conclusion
	// reference types
	// slice map channel pointer function
	// value types
	// int float string bool struct
}

// pass by value
// func (p person) updateName(newFirstName string) {
// 	p.firstName = newFirstName
// }

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func updateGreeting(greeting []string) {
	greeting[0] = "Bye"
}
