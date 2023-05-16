package main

import "fmt"

type PrintRequestHandler struct {
}

func NewPrintRequestHandler() *PrintRequestHandler {
	return &PrintRequestHandler{}
}

func (prh PrintRequestHandler) Handle(request *Request) {
	fmt.Println("print...")
	fmt.Println("default request handler")
	fmt.Println(request.Body())
}
