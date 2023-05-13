package main

type ShipFactoryDetail interface {
	CreateShip(name string) *Ship
}
