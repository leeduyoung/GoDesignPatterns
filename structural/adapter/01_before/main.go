package main

import "fmt"

func main() {
	var searchService SearchServiceInterface
	searchService = new(SearchService)

	fmt.Println(
		searchService.Search("design_pattern"),
	)
}
