package gumball_machine

import (
	"fmt"
	"github.com/leeduyoung/GoDesignPatterns/behavioral/state/04_after/gumball_interface"
	"github.com/leeduyoung/GoDesignPatterns/behavioral/state/04_after/state"
)

type gumballMachine struct {
	state gumball_interface.State
	count int

	soldOutState    gumball_interface.State
	soldState       gumball_interface.State
	noQuarterState  gumball_interface.State
	hasQuarterState gumball_interface.State
}

func NewGumballMachine(count int) gumball_interface.GumballMachine {
	gm := new(gumballMachine)
	gm.soldOutState = state.NewSoldOutState(gm)
	gm.soldState = state.NewSoldState(gm)
	gm.noQuarterState = state.NewNoQuarterState(gm)
	gm.hasQuarterState = state.NewHasQuarterState(gm)
	gm.count = count
	gm.state = gm.soldState

	if count > 0 {
		gm.state = gm.noQuarterState
	}

	return gm
}

func (gm gumballMachine) InsertQuarter() {
	gm.state.InsertQuarter()
}

func (gm gumballMachine) EjectQuarter() {
	gm.state.EjectQuarter()
}

func (gm gumballMachine) TurnCrank() {
	gm.state.TurnCrank()
	if gm.state == gm.soldState {
		gm.state.Dispense()
	}
}

func (gm *gumballMachine) ChangeState(state gumball_interface.State) {
	gm.state = state
}

func (gm *gumballMachine) ReleaseBall() {
	fmt.Println("캡슐을 내보내고 있습니다.")
	if gm.count > 0 {
		gm.count = gm.count - 1
	}
}

func (gm *gumballMachine) Count() int {
	return gm.count
}

func (gm *gumballMachine) State() gumball_interface.State {
	return gm.state
}

func (gm *gumballMachine) SoldOutState() gumball_interface.State {
	return gm.soldOutState
}

func (gm *gumballMachine) SoldState() gumball_interface.State {
	return gm.soldState
}

func (gm *gumballMachine) NoQuarterState() gumball_interface.State {
	return gm.noQuarterState
}

func (gm *gumballMachine) HasQuarterState() gumball_interface.State {
	return gm.hasQuarterState
}
