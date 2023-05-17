package main

type Book struct {
	Title  string
	Author string
}

func (b Book) Accept(v Visitor) {
	v.VisitBook(b)
}
