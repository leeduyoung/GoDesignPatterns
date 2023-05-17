package main

import "fmt"

func main() {
	// Create books and videos
	books := []Book{
		{Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"},
		{Title: "To Kill a Mockingbird", Author: "Harper Lee"},
	}

	videos := []Video{
		{Title: "The Shawshank Redemption", Director: "Frank Darabont", Description: "Two imprisoned men bond over several years."},
		{Title: "Pulp Fiction", Director: "Quentin Tarantino", Description: "Various interconnected stories of criminals."},
	}

	// Calculate total price
	totalPrice := CalculateTotalPrice(books, videos)
	fmt.Printf("Total price: $%.2f\n", totalPrice)

	// Count inventory
	bookCount, videoCount := CountInventory(books, videos)
	fmt.Println("Inventory:")
	fmt.Println("Books:", bookCount)
	fmt.Println("Videos:", videoCount)
}
