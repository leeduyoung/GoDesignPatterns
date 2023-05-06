package main

type FrontDesk struct {
	restaurant      *Restaurant
	cleaningService *CleaningService
}

func NewFrontDesk() *FrontDesk {
	return &FrontDesk{
		restaurant:      NewRestaurant(),
		cleaningService: NewCleaningService(),
	}
}

func (fd *FrontDesk) GetTower(guest Guest, number int) {
	// TODO: guest의 roomNumber 조회
	roomNumber := 1131

	fd.cleaningService.GetTower(roomNumber, number)
}

func (fd *FrontDesk) Dinner(guest Guest) {
	fd.restaurant.Dinner(guest)
}

func (fd *FrontDesk) Clean() {
	fd.cleaningService.Clean(*fd.restaurant)
}
