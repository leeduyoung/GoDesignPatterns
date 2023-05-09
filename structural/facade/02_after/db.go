package main

type Database struct {
	personMap map[string]*Person
}

func NewDatabase() *Database {
	return &Database{
		personMap: make(map[string]*Person),
	}
}

func (d *Database) Query(name string) *Person {
	if p, exists := d.personMap[name]; exists {
		return p
	}

	return nil
}
