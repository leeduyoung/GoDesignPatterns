package main

import "fmt"

func main() {

	prototype := NewPerson("John", 34, "A")

	person := prototype.Clone("Jane")
	person2 := prototype.Clone("Kaye")

	fmt.Println(prototype)
	fmt.Println(person)
	fmt.Println(person2)
}
