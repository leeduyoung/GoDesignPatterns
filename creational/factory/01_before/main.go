package main

import "fmt"

func main() {
	shipFactory := NewShipFactory()
	whiteShip := shipFactory.OrderShip("whiteShip", "ship@gmail.com")
	fmt.Println(whiteShip)

	blackShip := shipFactory.OrderShip("blackShip", "ship@gmail.com")
	fmt.Println(blackShip)
}
