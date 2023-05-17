package main

import "fmt"

// CalculateTotalPrice Calculate total price of books and videos
func CalculateTotalPrice(books []Book, videos []Video) float64 {
	totalPrice := 0.0

	for _, book := range books {
		fmt.Println(book)
		totalPrice += 10.99
	}

	for _, video := range videos {
		fmt.Println(video)
		totalPrice += 19.99
	}

	return totalPrice
}
