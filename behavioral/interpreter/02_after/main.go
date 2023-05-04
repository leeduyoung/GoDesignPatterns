package main

import "fmt"

func main() {
	postfixExpression := PostfixParserParse("xyz+-")
	result := postfixExpression.Interpret(map[string]int{
		"x": 1,
		"y": 2,
		"z": 3,
	})
	fmt.Println(result) // -4
}
