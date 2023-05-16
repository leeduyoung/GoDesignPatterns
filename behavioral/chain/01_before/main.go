package main

func main() {
	requestHandler := NewLoggingRequestHandler()
	requestHandler2 := NewPrintRequestHandler()
	requestHandler3 := NewLoggingAndPrintRequestHandler()

	request := NewRequest("https://ithub.tistory.com")
	requestHandler.Handle(request)
	requestHandler2.Handle(request)
	requestHandler3.Handle(request)
}
