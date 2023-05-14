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

func (p Person) Clone(name string) *Person {

	// TODO: 변경이 필요한 부분 처리

	return &Person{
		name:  name,
		age:   p.age,
		class: p.class,
	}
}
