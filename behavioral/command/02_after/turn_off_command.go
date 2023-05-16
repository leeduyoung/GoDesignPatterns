package main

type TurnOffCommand struct {
	light *Light
}

func NewTurnOffCommand(light *Light) *TurnOffCommand {
	return &TurnOffCommand{
		light: light,
	}
}

func (toc TurnOffCommand) Execute() {
	toc.light.TurnOff()
}
