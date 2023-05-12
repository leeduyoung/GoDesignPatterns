package main

import "fmt"

func main() {
	/**
	CASE 1.
	*/

	//// CREATE WHITE SHIP
	//whiteShip := NewShipFactory(NewWhiteShipFactory()).OrderShip("whiteShip", "kaye@gmail.com")
	//fmt.Println(whiteShip)
	//
	//// CREATE BLACK SHIP
	//blackShip := NewShipFactory(NewBlackShipFactory()).OrderShip("blackShip", "kaye@gmail.com")
	//fmt.Println(blackShip)

	/**
	CASE 2.
	*/

	// CREATE WHITE SHIP
	dip(NewWhiteShipFactory(), "whiteShip", "kaye@gmail.com")

	// CREATE BLACK SHIP
	dip(NewBlackShipFactory(), "blackShip", "kaye@gmail.com")

}

// dip 의존성 주입 예제
func dip(shipFactoryDetail ShipFactoryDetail, name, email string) {
	ship := NewShipFactory(shipFactoryDetail).OrderShip(name, email)
	fmt.Println(ship)
}
