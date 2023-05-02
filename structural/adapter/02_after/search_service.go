package main

var keywordMap = map[string][]string{
	"design_pattern": {"structural", "creational", "behavioral"},
	"structural":     {"composite", "proxy"},
	"creational":     {"builder", "factory"},
	"behavioral":     {"state", "template"},
}

type SearchService struct {
}

func (s *SearchService) Search(keyword string) []string {
	// DO SOMETHING...
	if v, exists := keywordMap[keyword]; exists {
		return v
	}

	return []string{}
}
