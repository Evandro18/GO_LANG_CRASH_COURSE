package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName, lastName, city, gender string
	age int
}

//Greeting method (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday()  {
	p.age++
}
// getMarried
func (p *Person) getMarried(spouseLastName string)  {
	if p.gender == "m" {
		return
	}
	p.lastName = spouseLastName
}

func main() {
	// Initialize person using struct
	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 25}
	person2 := Person{firstName: "Bob", lastName: "Johnson", city: "New York", gender: "m", age: 26}
	// person1 := Person{"Samantha", "Smith", "Boston", "f", 5}
	// fmt.Println(person1.firstName)
	// fmt.Println(person1.age)
	// person1.age++
	person1.hasBirthday()
	person1.hasBirthday()
	person1.hasBirthday()
	person1.hasBirthday()
	person1.getMarried("Willians")
	person2.getMarried("Thompson")
	fmt.Println(person2.greet())
}