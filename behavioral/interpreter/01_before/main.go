package main

import "fmt"

func main() {
	postfixNotation := NewPostfixNotation("123+-")
	result := postfixNotation.Calculate()

	fmt.Println(result) // -4
}
