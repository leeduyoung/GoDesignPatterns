package main

func main() {
	// 의존관계가 복잡하게 얽혀있다.
	cleaningService := NewCleaningService()

	restaurant := NewRestaurant(cleaningService)
	restaurant.Clean()

	guest := NewGuest(restaurant, cleaningService)
	guest.GetTower(3)
	guest.Dinner()
}
