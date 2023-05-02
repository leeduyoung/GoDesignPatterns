package main

type SearchServiceInterface interface {
	Search(keyword string) []string
}
