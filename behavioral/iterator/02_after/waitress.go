package main

import "fmt"

type Waitress struct {
	pancakeHouseMenu Menu
	dinnerMenu       Menu
}

func NewWaitress(pancakeHouseMenu Menu, dinnerMenu Menu) *Waitress {
	return &Waitress{
		pancakeHouseMenu: pancakeHouseMenu,
		dinnerMenu:       dinnerMenu,
	}
}

func (w *Waitress) PrintMenu() {
	pancakeHouseMenuIterator := w.pancakeHouseMenu.CreateIterator()
	printMenu(pancakeHouseMenuIterator)

	dinnerMenuIterator := w.dinnerMenu.CreateIterator()
	printMenu(dinnerMenuIterator)
}

func printMenu(iterator Iterator[MenuItem]) {
	for iterator.HasNext() {
		menuItem := iterator.Next()

		fmt.Println(menuItem.Name())
		fmt.Println(menuItem.Description())
		fmt.Println(menuItem.IsVegetarian())
		fmt.Println(menuItem.Price())
	}
}
