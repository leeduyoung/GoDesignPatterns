package main

func main() {
	// Create the receiver
	light := new(Light)

	// Create the command
	turnOnCommand := NewTurnOnCommand(light)
	turnOffCommand := NewTurnOffCommand(light)

	// Create the invoker
	remoteController := new(RemoteController)

	// Set the command and excute
	remoteController.SetOnCommand(turnOnCommand)
	remoteController.SetOffCommand(turnOffCommand)

	remoteController.PressOnButton()
	remoteController.PressOffButton()
}
