package main

type ChatServer struct {
	messageMap map[string][]string
}

func NewChatServer() *ChatServer {
	return &ChatServer{
		messageMap: make(map[string][]string),
	}
}

func (cs *ChatServer) AddMessage(subject, message string) {
	if v, exists := cs.messageMap[subject]; exists {
		cs.messageMap[subject] = append(v, message)
		return
	}

	cs.messageMap[subject] = []string{message}
}

func (cs *ChatServer) Messages(subject string) []string {
	if v, exists := cs.messageMap[subject]; exists {
		return v
	}

	return nil
}
