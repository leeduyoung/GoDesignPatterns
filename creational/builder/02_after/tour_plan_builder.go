package main

import "time"

type TourPlanBuilder interface {
	Title(title string) TourPlanBuilder
	StartDate(startDate time.Time) TourPlanBuilder
	NightsAndDays(nights, days int) TourPlanBuilder
	WhereToStay(whereToStay string) TourPlanBuilder
	AddPlan(day int, plan string) TourPlanBuilder
	GetTourPlan() *TourPlan
}
