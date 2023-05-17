package main

type Video struct {
	Title       string
	Director    string
	Description string
}

func (v Video) Accept(visitor Visitor) {
	visitor.VisitVideo(v)
}
