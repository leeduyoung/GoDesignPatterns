package gumball_interface

type GumballMachine interface {
	InsertQuarter()
	EjectQuarter()
	TurnCrank()
	ChangeState(state State)
	SoldOutState() State
	SoldState() State
	NoQuarterState() State
	HasQuarterState() State
	ReleaseBall()
	Count() int
	State() State
}
