package main

import (
	"fmt"
	"time"
)

func main() {
	// SHORT TRIP
	shortTrip := NewTourPlan()
	shortTrip.SetTitle("오레곤 롱비치 여행")
	shortTrip.SetStartDate(time.Date(2025, 12, 25, 13, 0, 0, 0, time.Local))

	fmt.Println(shortTrip)

	// LONG TRIP
	longTrip := NewTourPlan()
	longTrip.SetTitle("칸쿤 여행")
	longTrip.SetNights(4)
	longTrip.SetDays(5)
	longTrip.SetStartDate(time.Date(2025, 7, 25, 13, 0, 0, 0, time.Local))
	longTrip.SetWhereToStay("호텔")
	longTrip.AddPlan(0, "체크인 후 짐풀기")
	longTrip.AddPlan(0, "저녁 식사")
	longTrip.AddPlan(1, "조식 부페에서 식사")
	longTrip.AddPlan(1, "해변가 산책")
	longTrip.AddPlan(1, "호텔 수영장에서 놀기")
	longTrip.AddPlan(1, "호텔 근처 레스토랑에서 저녁식사")
	longTrip.AddPlan(2, "조식 부페에서 식사")
	longTrip.AddPlan(2, "낮잠")
	longTrip.AddPlan(2, "휴식")
	longTrip.AddPlan(2, "저녁식사")
	longTrip.AddPlan(2, "맥주 파티")

	fmt.Println(longTrip)
}
