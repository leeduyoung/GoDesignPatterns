package main

import "fmt"

type LoggingRequestHandler struct {
	nextHandler RequestHandlerInterface
}

func NewLoggingRequestHandler(handler RequestHandlerInterface) *LoggingRequestHandler {
	return &LoggingRequestHandler{
		nextHandler: handler,
	}
}

func (lrh LoggingRequestHandler) Handle(request *Request) {
	fmt.Println("logging...")
	lrh.nextHandler.Handle(request)
}
