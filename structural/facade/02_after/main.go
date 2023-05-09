package main

import "fmt"

func main() {
	storage := NewStorage()
	person := storage.Query("kaye")

	fmt.Println(person)
}
