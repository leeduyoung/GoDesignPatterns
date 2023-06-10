package state

import (
	"fmt"
	"github.com/leeduyoung/GoDesignPatterns/behavioral/state/04_after/gumball_interface"
)

type HasQuarterState struct {
	gumballMachine gumball_interface.GumballMachine
}

func (s HasQuarterState) Dispense() {
	fmt.Println("오류")
}

func (s HasQuarterState) InsertQuarter() {
	fmt.Println("동전은 한개만 넣어주세요.")
}

func (s HasQuarterState) EjectQuarter() {
	s.gumballMachine.ChangeState(s.gumballMachine.NoQuarterState())
	fmt.Println("동전이 반환되었습니다.")
}

func (s HasQuarterState) TurnCrank() {
	fmt.Println("손잡이를 돌리셨습니다.")
	s.gumballMachine.ChangeState(s.gumballMachine.SoldState())
}

func NewHasQuarterState(gm gumball_interface.GumballMachine) *HasQuarterState {
	return &HasQuarterState{
		gumballMachine: gm,
	}
}
