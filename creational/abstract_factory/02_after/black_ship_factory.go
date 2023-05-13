package main

type BlackShipFactory struct {
	ShipPartsFactory
}

func NewBlackShipFactory(partsFactory ShipPartsFactory) *BlackShipFactory {
	return &BlackShipFactory{
		partsFactory,
	}
}

func (bsf BlackShipFactory) CreateShip(name string) *Ship {
	var (
		color = "black"
		logo  = "blackShipLogo"
	)

	// CREATE WHEEL & ANCHOR  <- 이부분을 바꾸고 싶다면 어떻게 될까?
	wheel := bsf.CreateWheel()
	anchor := bsf.CreateAnchor()

	return NewShip(name, color, logo, wheel, anchor)
}
