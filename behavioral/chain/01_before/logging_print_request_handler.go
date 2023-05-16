package main

import "fmt"

type LoggingAndPrintRequestHandler struct {
}

func NewLoggingAndPrintRequestHandler() *LoggingAndPrintRequestHandler {
	return &LoggingAndPrintRequestHandler{}
}

func (laprh LoggingAndPrintRequestHandler) Handle(request *Request) {
	fmt.Println("logging & print...")
	fmt.Println("default request handler")
	fmt.Println(request.Body())
}
