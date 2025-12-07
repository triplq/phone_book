package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Contact struct {
	Person
	Number string
}

func AddContact(person Person, number string) *Contact {
	return &Contact{
		Person: person,
		Number: number,
	}
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

func (p Person) isAdult() bool {
	return p.Age >= 18
}

func (p *Person) BD() {
	p.Age++
}

func main() {
	var contacts []Contact
	contacts = append(contacts, *AddContact(*NewPerson("Tom", 24), "89040928902"))
	contacts[0].BD()
	fmt.Println(contacts)
}
