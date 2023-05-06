package main

import "fmt"

type Restaurant struct {
}

func NewRestaurant() *Restaurant {
	return &Restaurant{}
}

func (r Restaurant) Dinner(g Guest) {
	fmt.Println("[Restaurant] 식사를 한다.")
}
