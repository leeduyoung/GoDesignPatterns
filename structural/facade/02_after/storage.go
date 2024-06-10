package main

type Storage struct {
	db    *Database
	cache *Cache
}

func NewStorage() *Storage {
	return &Storage{
		db:    NewDatabase(),
		cache: NewCache(),
	}
}

func (*Storage) Query(name string) *Person {
	// 사용자 정보 조회

	// Cache 조회
	cache := NewCache()
	person := cache.Query("kaye")

	// Cache에 없다면 DB 조회
	if person == nil {
		db := NewDatabase()
		person = db.Query("kaye")
	}

	return person
}
