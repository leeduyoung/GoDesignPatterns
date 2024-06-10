package main

import "fmt"

func main() {
	// 사용자 정보 조회

	// Cache 조회
	cache := NewCache()
	person := cache.Query("kaye")

	// Cache에 없다면 DB 조회
	if person == nil {
		db := NewDatabase()
		person = db.Query("kaye")
	}

	fmt.Println(person)
}
