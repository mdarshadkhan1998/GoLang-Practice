package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}
type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	alex := person{"Alex", "Anderson", contactInfo{"arshad@gmail.com", 11111}}
	arshad := person{firstName: "Arshad", lastName: "Khan", contact: contactInfo{email: "arshad@gmail.com", zipCode: 1111}}
	fmt.Println(alex, arshad)
	var kh person
	fmt.Printf("%+v", kh)
	kh.firstName = "nax"
	kh.lastName = "kha"
	kh.contact.email = "arshad@gmail.com"
	fmt.Printf("%+v", kh)

	alex.print()

	arshPointer := &arshad
	arshPointer.updateName("MD Arshad")
}
func (personPointer *person) updateName(newFirstName string) {
	personPointer.firstName = newFirstName
	fmt.Println(personPointer)
	fmt.Println(*personPointer)
}

func (p person) print() {
	fmt.Printf("Inside print function %+v", p)
}
