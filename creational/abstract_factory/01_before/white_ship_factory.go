package main

import "fmt"

type WhiteShipFactory struct {
}

func NewWhiteShipFactory() *WhiteShipFactory {
	return &WhiteShipFactory{}
}

func (wsf WhiteShipFactory) CreateShip(name string) *Ship {
	var (
		color = "white"
		logo  = "whiteShipLogo"
	)

	// CREATE WHEEL & ANCHOR
	wheel := new(Wheel)
	anchor := new(Anchor)
	fmt.Println("set wheel & anchor")

	return NewShip(name, color, logo, *wheel, *anchor)
}
