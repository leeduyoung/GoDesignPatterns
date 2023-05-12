package main

type BlackShipFactory struct {
}

func NewBlackShipFactory() *BlackShipFactory {
	return &BlackShipFactory{}
}

func (wsf *BlackShipFactory) CreateShip(name string) *Ship {
	// Customizing for specific name
	var (
		color = "black"
		logo  = "blackShipLogo"
	)

	// Create ship
	return NewShip(name, color, logo)
}
