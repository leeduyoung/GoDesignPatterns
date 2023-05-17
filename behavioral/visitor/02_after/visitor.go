package main

type Visitor interface {
	VisitBook(book Book)
	VisitVideo(video Video)
}
