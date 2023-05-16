package main

type RemoteController struct {
	onCommand  Command
	offCommand Command
}

func NewRemoteController() *RemoteController {
	return &RemoteController{}
}

func (rc *RemoteController) SetOnCommand(command Command) {
	rc.onCommand = command
}

func (rc *RemoteController) SetOffCommand(command Command) {
	rc.offCommand = command
}

func (rc RemoteController) PressOnButton() {
	rc.onCommand.Execute()
}

func (rc RemoteController) PressOffButton() {
	rc.offCommand.Execute()
}
