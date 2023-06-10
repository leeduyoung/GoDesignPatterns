package state

import (
	"fmt"
	"github.com/leeduyoung/GoDesignPatterns/behavioral/state/04_after/gumball_interface"
)

type SoldState struct {
	gumballMachine gumball_interface.GumballMachine
}

func (s SoldState) Dispense() {
	s.gumballMachine.ReleaseBall()

	if s.gumballMachine.Count() > 0 {
		s.gumballMachine.ChangeState(s.gumballMachine.NoQuarterState())
	} else {
		fmt.Println("캡슐이 매진되었습니다.")
		s.gumballMachine.ChangeState(s.gumballMachine.SoldOutState())
	}
}

func (s SoldState) InsertQuarter() {
	fmt.Println("캡슐을 내보내고 있습니다.")
}

func (s SoldState) EjectQuarter() {
	fmt.Println("이미 캡슐이 반환되었습니다.")
}

func (s SoldState) TurnCrank() {
	fmt.Println("손잡이는 한번만 돌려주세요.")
}

func NewSoldState(gm gumball_interface.GumballMachine) *SoldState {
	return &SoldState{
		gumballMachine: gm,
	}
}
