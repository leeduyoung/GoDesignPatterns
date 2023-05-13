package main

import "fmt"

func main() {
	// CREATE WHITE SHIP
	whiteShip := NewShipFactory(
		NewWhiteShipFactory(new(BlackShipPartsFactory)),
	).OrderShip("whiteShip", "kaye@gmail.com")
	fmt.Println(whiteShip)

	// CREATE BLACK SHIP
	blackShip := NewShipFactory(
		NewBlackShipFactory(new(BlackShipPartsFactory)),
	).OrderShip("blackShip", "kaye@gmail.com")
	fmt.Println(blackShip)
}
