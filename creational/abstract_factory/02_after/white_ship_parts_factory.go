package main

import "fmt"

type WhiteShipPartsFactory struct {
}

func (wspf WhiteShipPartsFactory) CreateWheel() Wheel {
	fmt.Println("set white wheel")
	return new(WhiteWheel)
}

func (wspf WhiteShipPartsFactory) CreateAnchor() Anchor {
	fmt.Println("set white anchor")
	return new(WhiteAnchor)
}
