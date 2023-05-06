package main

func main() {
	chatServer := NewChatServer()
	user1 := NewUser("케이")
	user2 := NewUser("엘린")

	chatServer.Join("디자인 패턴", user1)
	chatServer.Join("디자인 패턴", user2)

	chatServer.SendMessage(user1, "디자인 패턴", "옵저버 패턴에 대해서 알아볼까요?")
	chatServer.SendMessage(user1, "디자인 패턴", "옵저버 패턴이란...")

	chatServer.Left("디자인 패턴", user2)

	chatServer.SendMessage(user1, "디자인 패턴", "이해가 안되나요..?")
}
