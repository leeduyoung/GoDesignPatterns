package main

func main() {
	requestHandler := NewLoggingRequestHandler(
		NewPrintRequestHandler(
			NewRequestHandler(),
		),
	)

	request := NewRequest("https://ithub.tistory.com")
	requestHandler.Handle(request)
}
