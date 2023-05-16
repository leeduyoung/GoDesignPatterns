package main

/**
커맨드 패턴을 적용하기 전 코드가 더 간결하다고 생각할 수 있지만
클라이언트와 수신자가 밀접하게 연결되어 있어서 새로운 명령을 도입하기 어렵다.
*/

func main() {
	// Create the receiver
	light := new(Light)

	// Turn on the light
	light.TurnOn()

	// Turn off the light
	light.TurnOff()
}
