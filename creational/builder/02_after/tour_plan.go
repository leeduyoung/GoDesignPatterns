package main

import "time"

type TourPlan struct {
	title       string
	nights      int
	days        int
	startDate   time.Time
	whereToStay string
	plans       []*DetailPlan
}

func (tp TourPlan) Title() string {
	return tp.title
}

func (tp *TourPlan) SetTitle(title string) {
	tp.title = title
}

func (tp TourPlan) Nights() int {
	return tp.nights
}

func (tp *TourPlan) SetNights(nights int) {
	tp.nights = nights
}

func (tp TourPlan) Days() int {
	return tp.days
}

func (tp *TourPlan) SetDays(days int) {
	tp.days = days
}

func (tp TourPlan) StartDate() time.Time {
	return tp.startDate
}

func (tp *TourPlan) SetStartDate(startDate time.Time) {
	tp.startDate = startDate
}

func (tp TourPlan) WhereToStay() string {
	return tp.whereToStay
}

func (tp *TourPlan) SetWhereToStay(whereToStay string) {
	tp.whereToStay = whereToStay
}

func (tp TourPlan) Plans() []*DetailPlan {
	return tp.plans
}

func (tp *TourPlan) AddPlan(day int, plan string) {
	tp.plans = append(tp.plans, NewDetailPlan(day, plan))
}

func (tp *TourPlan) SetPlans(plans []*DetailPlan) {
	tp.plans = plans
}
