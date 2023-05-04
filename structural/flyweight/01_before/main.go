package main

import "fmt"

func main() {
	characterA := NewCharacter("A", "black", "nanum", 16)
	characterB := NewCharacter("B", "black", "nanum", 16)
	characterC := NewCharacter("C", "black", "nanum", 16)
	characterD := NewCharacter("D", "black", "nanum", 16)
	characterE := NewCharacter("E", "black", "nanum", 16)
	characterF := NewCharacter("F", "black", "nanum", 16)
	characterG := NewCharacter("G", "black", "nanum", 16)

	fmt.Println(characterA)
	fmt.Println(characterB)
	fmt.Println(characterC)
	fmt.Println(characterD)
	fmt.Println(characterE)
	fmt.Println(characterF)
	fmt.Println(characterG)
}
