package main

import "fmt"

type RequestHandlerInterface interface {
	Handle(request *Request)
}

type RequestHandler struct {
}

func NewRequestHandler() *RequestHandler {
	return &RequestHandler{}
}

func (rh RequestHandler) Handle(request *Request) {
	fmt.Println("default request handler")
	fmt.Println(request.Body())
}
