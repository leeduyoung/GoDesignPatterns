package main

type MenuItem struct {
	name        string
	description string
	vegetarian  bool
	price       float64
}

func NewMenuItem(name, description string, vegetarian bool, price float64) *MenuItem {
	return &MenuItem{
		name:        name,
		description: description,
		vegetarian:  vegetarian,
		price:       price,
	}
}

func (m *MenuItem) Name() string {
	return m.name
}

func (m *MenuItem) Description() string {
	return m.description
}

func (m *MenuItem) Price() float64 {
	return m.price
}

func (m *MenuItem) IsVegetarian() bool {
	return m.vegetarian
}
