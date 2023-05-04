package main

type Character struct {
	value      string
	color      string
	fontFamily string
	fontSize   int
}

func NewCharacter(value, color, fontFamily string, fontSize int) *Character {
	return &Character{
		value:      value,
		color:      color,
		fontFamily: fontFamily,
		fontSize:   fontSize,
	}
}
