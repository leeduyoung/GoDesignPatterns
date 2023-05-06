package main

type Guest struct {
	restaurant      *Restaurant
	cleaningService *CleaningService
}

func NewGuest(restaurant *Restaurant, cleaningService *CleaningService) *Guest {
	return &Guest{
		restaurant:      restaurant,
		cleaningService: cleaningService,
	}
}

func (g *Guest) GetTower(number int) {
	g.cleaningService.GetTower(g, number)
}

func (g *Guest) Dinner() {
	g.restaurant.Dinner(g)
}
