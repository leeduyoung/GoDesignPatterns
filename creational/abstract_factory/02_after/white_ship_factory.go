package main

type WhiteShipFactory struct {
	ShipPartsFactory
}

func NewWhiteShipFactory(partsFactory ShipPartsFactory) *WhiteShipFactory {
	return &WhiteShipFactory{
		ShipPartsFactory: partsFactory,
	}
}

func (wsf WhiteShipFactory) CreateShip(name string) *Ship {
	var (
		color = "white"
		logo  = "whiteShipLogo"
	)

	// CREATE WHEEL & ANCHOR
	wheel := wsf.CreateWheel()
	anchor := wsf.CreateAnchor()

	return NewShip(name, color, logo, wheel, anchor)
}
