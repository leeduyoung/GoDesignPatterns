package main

import "fmt"

func main() {
	chatServer := NewChatServer()

	user1 := NewUser("케이", chatServer)
	user1.SendMessage("디자인 패턴", "옵저버 패턴에 대해서 알아봅시다.")
	user1.SendMessage("디자인 패턴", "옵저버 패턴이란...")

	// user2는 주기적으로 메시지를 조회해서 메시지를 확인해야한다.
	// 메시지가 정확히 언제 발송되었는지 알 수 없다.
	user2 := NewUser("엘린", chatServer)
	messages := user2.Messages("디자인 패턴")
	fmt.Println(messages)

	user1.SendMessage("디자인 패턴", "이해가 되세요..?")
}
