package gumball_interface

type State interface {
	InsertQuarter()
	EjectQuarter()
	TurnCrank()
	Dispense()
}
