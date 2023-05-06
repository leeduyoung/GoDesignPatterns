package main

type User struct {
	name       string
	chatServer *ChatServer
}

func NewUser(name string, chatServer *ChatServer) *User {
	return &User{
		name:       name,
		chatServer: chatServer,
	}
}

func (u *User) SendMessage(subject, message string) {
	userMessage := u.name + ": " + message
	u.chatServer.AddMessage(subject, userMessage)
}

func (u *User) Messages(subject string) []string {
	return u.chatServer.Messages(subject)
}
