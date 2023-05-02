package main

import (
	"fmt"
	"github.com/leeduyoung/GoDesignPatterns/structural/adapter/02_after/third_party"
)

func main() {
	var (
		searchService SearchServiceInterface
	)

	queryService := new(third_party.Query)
	searchService = NewSearchServiceAdapter(queryService)

	fmt.Println(
		searchService.Search("design_pattern"),
	)
}
