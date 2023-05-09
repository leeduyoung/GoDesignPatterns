package main

type Person struct {
	name   string
	age    int
	height int
}

func NewPerson(name string, age, height int) *Person {
	return &Person{
		name:   name,
		age:    age,
		height: height,
	}
}

func (p *Person) Name() string {
	return p.name
}

func (p *Person) Age() int {
	return p.age
}

func (p *Person) Height() int {
	return p.height
}
