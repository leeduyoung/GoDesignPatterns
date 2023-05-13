package main

import "time"

type DefaultTourPlanBuilder struct {
	title       string
	nights      int
	days        int
	startDate   time.Time
	whereToStay string
	plans       []*DetailPlan
}

func NewDefaultTourPlanBuilder() *DefaultTourPlanBuilder {
	return &DefaultTourPlanBuilder{}
}

func (dtpb *DefaultTourPlanBuilder) Title(title string) TourPlanBuilder {
	dtpb.title = title
	return dtpb
}

func (dtpb *DefaultTourPlanBuilder) StartDate(startDate time.Time) TourPlanBuilder {
	dtpb.startDate = startDate
	return dtpb
}

func (dtpb *DefaultTourPlanBuilder) NightsAndDays(nights, days int) TourPlanBuilder {
	dtpb.nights = nights
	dtpb.days = days
	return dtpb
}

func (dtpb *DefaultTourPlanBuilder) WhereToStay(whereToStay string) TourPlanBuilder {
	dtpb.whereToStay = whereToStay
	return dtpb
}

func (dtpb *DefaultTourPlanBuilder) AddPlan(day int, plan string) TourPlanBuilder {
	dtpb.plans = append(dtpb.plans, NewDetailPlan(day, plan))
	return dtpb
}

func (dtpb *DefaultTourPlanBuilder) GetTourPlan() *TourPlan {
	tp := NewTourPlan()
	tp.SetTitle(dtpb.title)
	tp.SetNights(dtpb.nights)
	tp.SetDays(dtpb.days)
	tp.SetStartDate(dtpb.startDate)
	tp.SetWhereToStay(dtpb.whereToStay)
	tp.SetPlans(dtpb.plans)
	return tp
}
