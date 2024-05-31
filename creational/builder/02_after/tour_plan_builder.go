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

type tourPlanBuilder struct {
	title       string
	nights      int
	days        int
	startDate   time.Time
	whereToStay string
	plans       []*DetailPlan
}

func (t tourPlanBuilder) Title(title string) TourPlanBuilder {
	t.title = title
	return t
}

func (t tourPlanBuilder) StartDate(startDate time.Time) TourPlanBuilder {
	t.startDate = startDate
	return t
}

func (t tourPlanBuilder) WhereToStay(whereToStay string) TourPlanBuilder {
	t.whereToStay = whereToStay
	return t
}

func (t tourPlanBuilder) NightsAndDays(nights, days int) TourPlanBuilder {
	t.nights = nights
	t.days = days
	return t
}

func (t tourPlanBuilder) AddPlan(day int, plan string) TourPlanBuilder {
	t.plans = append(t.plans, NewDetailPlan(day, plan))
	return t
}

func (t tourPlanBuilder) GetTourPlan() *TourPlan {
	return &TourPlan{
		title:       t.title,
		nights:      t.nights,
		days:        t.days,
		startDate:   t.startDate,
		whereToStay: t.whereToStay,
		plans:       t.plans,
	}
}

func NewTourPlanBuilder() TourPlanBuilder {
	return &tourPlanBuilder{}
}
