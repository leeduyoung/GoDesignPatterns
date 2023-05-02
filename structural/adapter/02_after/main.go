package main

import (
	"fmt"
	"github.com/leeduyoung/GoDesignPatterns/structural/adapter/02_after/third_party"
)

func main() {
	queryService := new(third_party.Query)
	var searchService SearchServiceInterface = NewSearchServiceAdapter(queryService)
	
	fmt.Println(
		searchService.Search("design_pattern"),
	)
}
