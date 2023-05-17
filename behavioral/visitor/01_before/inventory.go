package main

// CountInventory Count the number of books and videos
func CountInventory(books []Book, videos []Video) (int, int) {
	bookCount := len(books)
	videoCount := len(videos)

	return bookCount, videoCount
}
