package main

type VariableExpression struct {
	char string
}

func NewVariableExpression(char string) *VariableExpression {
	return &VariableExpression{
		char: char,
	}
}

func (ve *VariableExpression) Interpret(context map[string]int) int {
	return context[ve.char]
}
