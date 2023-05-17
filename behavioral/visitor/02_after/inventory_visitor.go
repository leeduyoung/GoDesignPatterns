package main

type InventoryVisitor struct {
	Books  int
	Videos int
}

func (iv *InventoryVisitor) VisitBook(b Book) {
	iv.Books++
}

func (iv *InventoryVisitor) VisitVideo(v Video) {
	iv.Videos++
}
