package main

import "fmt"

func main() {
	font := NewFontFactory().getInstance("nanum:16")
	characterA := NewCharacter("A", "black", font.fontFamily, font.fontSize)
	characterB := NewCharacter("B", "black", font.fontFamily, font.fontSize)
	characterC := NewCharacter("C", "black", font.fontFamily, font.fontSize)
	characterD := NewCharacter("D", "black", font.fontFamily, font.fontSize)
	characterE := NewCharacter("E", "black", font.fontFamily, font.fontSize)
	characterF := NewCharacter("F", "black", font.fontFamily, font.fontSize)
	characterG := NewCharacter("G", "black", font.fontFamily, font.fontSize)

	fmt.Println(characterA)
	fmt.Println(characterB)
	fmt.Println(characterC)
	fmt.Println(characterD)
	fmt.Println(characterE)
	fmt.Println(characterF)
	fmt.Println(characterG)
}
