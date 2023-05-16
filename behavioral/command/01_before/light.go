package main

import "fmt"

type Light struct {
	isOn bool
}

func (l *Light) TurnOn() {
	l.isOn = true
	fmt.Println("Light turn on. :)")
}

func (l *Light) TurnOff() {
	l.isOn = false
	fmt.Println("Light turn off. :(")
}
