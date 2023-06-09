package main

import "fmt"

type Waitress struct {
	pancakeHouseMenu *PancakeHouseMenu
	dinnerMenu       *DinnerMenu
}

func NewWaitress(pancakeHouseMenu *PancakeHouseMenu, dinnerMenu *DinnerMenu) *Waitress {
	return &Waitress{
		pancakeHouseMenu: pancakeHouseMenu,
		dinnerMenu:       dinnerMenu,
	}
}

func (w *Waitress) PrintMenu() {
	pItems := w.pancakeHouseMenu.MenuItems()
	for _, v := range pItems {
		fmt.Println(v.Name())
		fmt.Println(v.Description())
		fmt.Println(v.IsVegetarian())
		fmt.Println(v.Price())
	}

	fmt.Println("============================================")

	dItems := w.dinnerMenu.MenuItems()
	for _, v := range dItems {
		fmt.Println(v.Name())
		fmt.Println(v.Description())
		fmt.Println(v.IsVegetarian())
		fmt.Println(v.Price())
	}
}
