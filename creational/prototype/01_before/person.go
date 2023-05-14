package main

type Person struct {
	name  string
	age   int
	class string
}

func NewPerson(name string, age int, class string) *Person {
	return &Person{
		name:  name,
		age:   age,
		class: class,
	}
}
