package main

import "fmt"

func main() {
	// CREATE WHITE SHIP
	whiteShip := NewShipFactory(NewWhiteShipFactory()).OrderShip("whiteShip", "kaye@gmail.com")
	fmt.Println(whiteShip)

	// CREATE BLACK SHIP
	blackShip := NewShipFactory(NewBlackShipFactory()).OrderShip("blackShip", "kaye@gmail.com")
	fmt.Println(blackShip)
}
