package main

import "github.com/leeduyoung/GoDesignPatterns/behavioral/interpreter/02_after/stack"

func PostfixParserParse(expression string) PostfixExpression {
	newStack := stack.New[PostfixExpression]()
	for _, v := range expression {
		char := string(v)
		newStack.Push(getExpression(char, newStack))
	}

	return newStack.Pop()
}

func getExpression(char string, stack stack.Stack[PostfixExpression]) PostfixExpression {
	switch char {
	case "+":
		return NewPlusExpression(stack.Pop(), stack.Pop())
	case "-":
		right := stack.Pop()
		left := stack.Pop()
		return NewMinusExpression(right, left)
	default:
		return NewVariableExpression(char)
	}

	return nil
}
