package main

import "fmt"

type PrintRequestHandler struct {
	nextHandler RequestHandlerInterface
}

func NewPrintRequestHandler(handler RequestHandlerInterface) *PrintRequestHandler {
	return &PrintRequestHandler{
		nextHandler: handler,
	}
}

func (prh PrintRequestHandler) Handle(request *Request) {
	fmt.Println("print...")
	prh.nextHandler.Handle(request)
}
