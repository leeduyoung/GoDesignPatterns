package main

type PlusExpression struct {
	right PostfixExpression
	left  PostfixExpression
}

func NewPlusExpression(right, left PostfixExpression) *PlusExpression {
	return &PlusExpression{
		right: right,
		left:  left,
	}
}

func (pe *PlusExpression) Interpret(context map[string]int) int {
	return pe.left.Interpret(context) + pe.right.Interpret(context)
}
