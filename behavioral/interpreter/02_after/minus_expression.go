package main

type MinusExpression struct {
	right PostfixExpression
	left  PostfixExpression
}

func NewMinusExpression(right, left PostfixExpression) *MinusExpression {
	return &MinusExpression{
		right: right,
		left:  left,
	}
}

func (pe *MinusExpression) Interpret(context map[string]int) int {
	return pe.left.Interpret(context) - pe.right.Interpret(context)
}
