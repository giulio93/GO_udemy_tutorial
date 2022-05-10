package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	//contact   contactInfo
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Pary",
		contactInfo: contactInfo{
			email:   "j@com",
			zipCode: 94000,
		},
	}

	jim.UpdateName("Marco")
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	//fmt.Printf("%+v", jim)
}

func (p person) PrintInfo() {
	fmt.Printf("%+v", p)
}
func (pointerToPerson *person) UpdateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName

}
