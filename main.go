package main

import (
	types "restAPI/src/types"
)

func changeValue(p types.Person) {
	p.Name = "Doe"
	p.Age = 40
}

func main() {
	person := types.Person{Name: "John", Age: 30}
	changeValue(person)
	println(person.Name, person.Age)
}
