package main

import (
	"errors"
	"fmt"
)

type ShipFactory interface {
	OrderShip(name, email string) *Ship
}

type shipFactory struct {
	ShipFactoryDetail
}

func NewShipFactory(detail ShipFactoryDetail) ShipFactory {
	return &shipFactory{
		ShipFactoryDetail: detail,
	}
}

func (sf shipFactory) OrderShip(name, email string) *Ship {
	// Validate
	err := validate(name, email)
	if err != nil {
		return nil
	}

	// Prepare
	prepareFor(name)

	// Customizing for specific name
	ship := sf.CreateShip(name)

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
