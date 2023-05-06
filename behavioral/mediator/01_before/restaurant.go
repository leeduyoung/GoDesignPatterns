package main

import "fmt"

type Restaurant struct {
	cleaningService *CleaningService
}

func (r Restaurant) Dinner(g *Guest) {
	fmt.Println("[Restaurant] 식사를 한다.")
}

func (r Restaurant) Clean() {
	r.cleaningService.Clean(r)
}

func NewRestaurant(cleaningService *CleaningService) *Restaurant {
	return &Restaurant{
		cleaningService: cleaningService,
	}
}
