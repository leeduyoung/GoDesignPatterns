package main

type Guest struct {
	frontDesk *FrontDesk
}

func NewGuest(frontDesk *FrontDesk) *Guest {
	return &Guest{
		frontDesk: frontDesk,
	}
}

func (g Guest) GetTower(number int) {
	g.frontDesk.GetTower(g, number)
}

func (g Guest) Dinner() {
	g.frontDesk.Dinner(g)
}
