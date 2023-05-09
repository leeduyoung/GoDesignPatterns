package main

type Cache struct {
	personMap map[string]*Person
}

func NewCache() *Cache {
	m := make(map[string]*Person)
	m["kaye"] = &Person{
		name:   "kaye",
		age:    34,
		height: 174,
	}

	return &Cache{
		personMap: m,
	}
}

func (c *Cache) Query(name string) *Person {
	if p, exists := c.personMap[name]; exists {
		return p
	}

	return nil
}
