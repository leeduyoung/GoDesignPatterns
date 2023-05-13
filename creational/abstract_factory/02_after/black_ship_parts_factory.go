package main

import "fmt"

type BlackShipPartsFactory struct {
}

func (bspf BlackShipPartsFactory) CreateWheel() Wheel {
	fmt.Println("set black wheel")
	return new(WhiteWheel)
}

func (bspf BlackShipPartsFactory) CreateAnchor() Anchor {
	fmt.Println("set black anchor")
	return new(WhiteAnchor)
}
