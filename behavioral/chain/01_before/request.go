package main

type Request struct {
	body string
}

func NewRequest(body string) *Request {
	return &Request{
		body: body,
	}
}

func (r Request) Body() string {
	return r.body
}
