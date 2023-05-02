package main

import "github.com/leeduyoung/GoDesignPatterns/structural/adapter/02_after/third_party"

type SearchServiceAdapter struct {
	queryService third_party.QueryInterface
}

func NewSearchServiceAdapter(queryService third_party.QueryInterface) *SearchServiceAdapter {
	return &SearchServiceAdapter{
		queryService: queryService,
	}
}

func (ssa *SearchServiceAdapter) Search(keyword string) []string {
	return ssa.queryService.Query(keyword)
}
