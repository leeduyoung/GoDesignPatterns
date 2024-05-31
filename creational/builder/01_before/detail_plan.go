package main

type DetailPlan struct {
	day  int
	plan string
}

func NewDetailPlan(day int, plan string) *DetailPlan {
	return &DetailPlan{
		day:  day,
		plan: plan,
	}
}

func (dp DetailPlan) Day() int {
	return dp.day
}

func (dp *DetailPlan) SetDay(day int) {
	dp.day = day
}

func (dp DetailPlan) Plan() string {
	return dp.plan
}

func (dp *DetailPlan) SetPlan(plan string) {
	dp.plan = plan
}
