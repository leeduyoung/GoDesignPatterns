package main

import "fmt"

func main() {
	// 사용자 정보 조회

	// Cache 조회
	db := NewDatabase()
	person := db.Query("kaye")

	// Cache에 없다면 DB 조회
	if person == nil {
		cache := NewCache()
		person = cache.Query("kaye")
	}

	fmt.Println(person)
}
