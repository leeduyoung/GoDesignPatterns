package main

import "fmt"

type LoggingRequestHandler struct {
}

func NewLoggingRequestHandler() *LoggingRequestHandler {
	return &LoggingRequestHandler{}
}

func (lrh LoggingRequestHandler) Handle(request *Request) {
	fmt.Println("logging...")
	fmt.Println("default request handler")
	fmt.Println(request.Body())
}
