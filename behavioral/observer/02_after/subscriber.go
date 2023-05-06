package main

type Subscriber interface {
	HandleMessage(message string)
}
