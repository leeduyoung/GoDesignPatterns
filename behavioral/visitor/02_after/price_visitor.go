package main

type PriceVisitor struct {
	TotalPrice float64
}

func (pv *PriceVisitor) VisitBook(b Book) {
	pv.TotalPrice += 10.99
}

func (pv *PriceVisitor) VisitVideo(v Video) {
	pv.TotalPrice += 19.99
}
