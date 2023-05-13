package main

import "time"

type TourDirector struct {
	TourPlanBuilder
}

func NewTourDirector(builder TourPlanBuilder) *TourDirector {
	return &TourDirector{
		builder,
	}
}

func (td *TourDirector) CancunTrip() *TourPlan {
	return td.Title("칸쿤 여행").
		NightsAndDays(4, 5).
		StartDate(time.Date(2024, 12, 25, 6, 0, 0, 0, time.Local)).
		WhereToStay("호텔").
		AddPlan(0, "놀기").
		AddPlan(1, "먹기").
		AddPlan(2, "마시기").
		GetTourPlan()
}

func (td *TourDirector) LongBeachTrip() *TourPlan {
	return td.Title("롱비치 여행").
		StartDate(time.Date(2024, 8, 15, 6, 0, 0, 0, time.Local)).
		GetTourPlan()
}
