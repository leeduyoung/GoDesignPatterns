package main

func main() {
	pancakeHouseMenu := NewPancakeHouseMenu()
	dinnerMenu := NewDinnerMenu()

	waitress := NewWaitress(
		pancakeHouseMenu,
		dinnerMenu,
	)

	waitress.PrintMenu()
}
