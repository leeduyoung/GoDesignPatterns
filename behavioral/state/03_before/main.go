package main

import "fmt"

func main() {
	gumballMachine := NewGumballMachine(5)
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
