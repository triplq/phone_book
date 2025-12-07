package main

import (
	"fmt"
	"strconv"
)

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

func FindContact(contacts []Contact, name string) (int, *Contact) {
	for i, val := range contacts {
		if val.Name == name {
			return i, &val
		}
	}
	return -1, nil
}

func DeleteContact(contacts []Contact, name string) []Contact {
	ind, _ := FindContact(contacts, name)
	new_cont := make([]Contact, 0, len(contacts)-1)
	new_cont = append(new_cont, contacts[:ind]...)
	new_cont = append(new_cont, contacts[ind+1:]...)
	return new_cont
}

func PrintAll(contacts []Contact) {
	for _, val := range contacts {
		fmt.Println("Name: " + val.Name + " Age: " + strconv.Itoa(val.Age) + " Number: " + val.Number)
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
	contacts = append(contacts, *AddContact(*NewPerson("Alice", 12), "4123434"))
	contacts = append(contacts, *AddContact(*NewPerson("Sasha", 98), "3434343"))
	contacts = append(contacts, *AddContact(*NewPerson("Alex", 4), "4111040"))

	ind, p_Alex := FindContact(contacts, "Alex")
	fmt.Println(ind, p_Alex)

	contacts = DeleteContact(contacts, "Alex")

	PrintAll(contacts)
}
