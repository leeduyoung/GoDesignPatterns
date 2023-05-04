package main

import (
	"github.com/leeduyoung/GoDesignPatterns/behavioral/interpreter/01_before/stack"
	"strconv"
)

type PostfixNotation struct {
	expression string
}

func NewPostfixNotation(expression string) *PostfixNotation {
	return &PostfixNotation{
		expression: expression,
	}
}

func (pn *PostfixNotation) Calculate() int {
	newStack := stack.New[int]()

	for _, v := range pn.expression {
		char := string(v)

		switch char {
		case "+":
			newStack.Push(newStack.Pop() + newStack.Pop())
		case "-":
			right := newStack.Pop()
			left := newStack.Pop()
			newStack.Push(left - right)
		default:
			value, _ := strconv.Atoi(char)
			newStack.Push(value)
		}
	}

	return newStack.Pop()
}
