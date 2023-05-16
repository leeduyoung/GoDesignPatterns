package main

type TurnOnCommand struct {
	light *Light
}

func NewTurnOnCommand(light *Light) *TurnOnCommand {
	return &TurnOnCommand{
		light: light,
	}
}

func (toc TurnOnCommand) Execute() {
	toc.light.TurnOn()
}
