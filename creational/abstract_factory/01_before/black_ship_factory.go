package main

import "fmt"

type BlackShipFactory struct {
}

func NewBlackShipFactory() *BlackShipFactory {
	return &BlackShipFactory{}
}

func (wsf BlackShipFactory) CreateShip(name string) *Ship {
	var (
		color = "black"
		logo  = "blackShipLogo"
	)

	// CREATE WHEEL & ANCHOR  <- 이부분을 바꾸고 싶다면 어떻게 될까?
	wheel := new(Wheel)
	anchor := new(Anchor)
	fmt.Println("set wheel & anchor")

	return NewShip(name, color, logo, *wheel, *anchor)
}
