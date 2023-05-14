package main

import "fmt"

func main() {
	person1 := NewPerson("John", 34, "A")
	person2 := NewPerson("Jane", 23, "A")

	fmt.Println(person1)
	fmt.Println(person2)
}
