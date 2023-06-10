package state

import (
	"fmt"
	"github.com/leeduyoung/GoDesignPatterns/behavioral/state/04_after/gumball_interface"
)

type SoldOutState struct {
	gumballMachine gumball_interface.GumballMachine
}

func (s SoldOutState) Dispense() {
	fmt.Println("오류")
}

func (s SoldOutState) InsertQuarter() {
	fmt.Println("매진 되었습니다. 다음에 이용해 주세요.")
}

func (s SoldOutState) EjectQuarter() {
	fmt.Println("매진 되어 동전을 넣을 수 없습니다.")
}

func (s SoldOutState) TurnCrank() {
	fmt.Println("매진입니다.")
}

func NewSoldOutState(gm gumball_interface.GumballMachine) *SoldOutState {
	return &SoldOutState{
		gumballMachine: gm,
	}
}
