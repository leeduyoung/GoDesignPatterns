package main

type PostfixExpression interface {
	Interpret(context map[string]int) int
}
