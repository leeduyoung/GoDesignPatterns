package main

type ChatServer struct {
	subscriberMap map[string][]Subscriber
}

func NewChatServer() *ChatServer {
	return &ChatServer{
		subscriberMap: make(map[string][]Subscriber),
	}
}

func (cs *ChatServer) Join(subject string, subscriber Subscriber) {
	if v, exists := cs.subscriberMap[subject]; exists {
		cs.subscriberMap[subject] = append(v, subscriber)
		return
	}

	cs.subscriberMap[subject] = []Subscriber{subscriber}
}

func (cs *ChatServer) Left(subject string, subscriber Subscriber) {
	if v, exists := cs.subscriberMap[subject]; exists {
		removeIndex := getIndex(v, subscriber)
		cs.subscriberMap[subject] = append(v[:removeIndex], v[removeIndex+1:]...)
	}
}

func getIndex(subscribers []Subscriber, user Subscriber) int {
	for i, v := range subscribers {
		if v == user {
			return i
		}
	}

	// error!
	return 0
}

func (cs *ChatServer) SendMessage(user *User, subject, message string) {
	if v, exists := cs.subscriberMap[subject]; exists {
		userMessage := user.Name() + ": " + message

		for _, subscriber := range v {
			subscriber.HandleMessage(userMessage)
		}
	}
}
