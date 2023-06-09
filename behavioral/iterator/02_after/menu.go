package main

type Menu interface {
	CreateIterator() Iterator[MenuItem]
}
