package state

import (
	"fmt"
	"github.com/leeduyoung/GoDesignPatterns/behavioral/state/04_after/gumball_interface"
)

type NoQuarterState struct {
	gumballMachine gumball_interface.GumballMachine
}

func (s NoQuarterState) Dispense() {
	fmt.Println("오류")
}

func (s NoQuarterState) InsertQuarter() {
	s.gumballMachine.ChangeState(s.gumballMachine.HasQuarterState())
	fmt.Println("동전을 넣으셨습니다.")
}

func (s NoQuarterState) EjectQuarter() {
	fmt.Println("반환할 동전이 없습니다.")
}

func (s NoQuarterState) TurnCrank() {
	fmt.Println("먼저 동전을 넣어주세요.")
}

func NewNoQuarterState(gm gumball_interface.GumballMachine) *NoQuarterState {
	return &NoQuarterState{
		gumballMachine: gm,
	}
}
