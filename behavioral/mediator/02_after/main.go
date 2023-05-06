package main

func main() {
	frontDesk := NewFrontDesk()
	frontDesk.Clean()

	guest := NewGuest(frontDesk)
	guest.GetTower(3)
	guest.Dinner()

}
