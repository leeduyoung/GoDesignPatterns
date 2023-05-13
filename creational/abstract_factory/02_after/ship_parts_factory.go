package main

type ShipPartsFactory interface {
	CreateAnchor() Anchor
	CreateWheel() Wheel
}
