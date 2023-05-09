package main

type Person struct {
	name   string
	age    int
	height int
}

type Database struct {
	personMap map[string]Person
}

func NewDatabase() *Database {
	return &Database{}
}
