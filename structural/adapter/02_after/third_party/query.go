package third_party

type QueryInterface interface {
	Query(keyword string) []string
}
