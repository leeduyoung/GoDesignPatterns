package main

import "fmt"

func main() {
	// Create books and videos
	book1 := Book{Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"}
	book2 := Book{Title: "To Kill a Mockingbird", Author: "Harper Lee"}

	video1 := Video{Title: "The Shawshank Redemption", Director: "Frank Darabont", Description: "Two imprisoned men bond over several years."}
	video2 := Video{Title: "Pulp Fiction", Director: "Quentin Tarantino", Description: "Various interconnected stories of criminals."}

	// Create visitors
	priceVisitor := &PriceVisitor{}
	inventoryVisitor := &InventoryVisitor{}

	// Accept visitors
	book1.Accept(priceVisitor)
	book2.Accept(priceVisitor)

	video1.Accept(priceVisitor)
	video2.Accept(priceVisitor)

	book1.Accept(inventoryVisitor)
	book2.Accept(inventoryVisitor)

	video1.Accept(inventoryVisitor)
	video2.Accept(inventoryVisitor)

	// Print results
	fmt.Printf("Total price: $%.2f\n", priceVisitor.TotalPrice)
	fmt.Println("Inventory:")
	fmt.Println("Books:", inventoryVisitor.Books)
	fmt.Println("Videos:", inventoryVisitor.Videos)
}
