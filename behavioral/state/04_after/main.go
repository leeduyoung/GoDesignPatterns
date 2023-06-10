package main

import (
	"fmt"
	"github.com/leeduyoung/GoDesignPatterns/behavioral/state/04_after/gumball_machine"
)

func main() {
	gumballMachine := gumball_machine.NewGumballMachine(5)
	gumballMachine.InsertQuarter()
	gumballMachine.EjectQuarter()

	fmt.Println(gumballMachine)
	fmt.Println("===================================")

	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()
	gumballMachine.TurnCrank()

	fmt.Println(gumballMachine)
	fmt.Println("===================================")

	gumballMachine.EjectQuarter()
	gumballMachine.InsertQuarter()
	gumballMachine.EjectQuarter()

	fmt.Println(gumballMachine)
}
