package main

import (
	"errors"
	"fmt"
)

type ShipFactory struct {
}

func NewShipFactory() *ShipFactory {
	return &ShipFactory{}
}

func (sf *ShipFactory) OrderShip(name, email string) *Ship {
	// Validate
	err := validate(name, email)
	if err != nil {
		return nil
	}

	// Prepare
	prepareFor(name)

	// Customizing for specific name
	var (
		color string
		logo  string
	)
	if name == "whiteShip" {
		logo = "whiteShipLogo"
		color = "white"
	} else if name == "blackShip" {
		logo = "blackShipLogo"
		color = "black"
	}

	// Create ship
	ship := NewShip(name, color, logo)

	// Notify
	sendEmailTo(email, ship)

	return ship
}

func sendEmailTo(email string, ship *Ship) {
	// TODO: email 전송
	fmt.Println(ship.Name() + " 생성 완료!")
}

func prepareFor(name string) {
	fmt.Println(name + " 만들 준비 중...")
}

func validate(name string, email string) error {
	if name == "" {
		return errors.New("failed to OrderShip. name is empty")
	}
	if email == "" {
		return errors.New("failed to OrderShip. email is empty")
	}

	return nil
}
