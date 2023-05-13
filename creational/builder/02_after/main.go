package main

func main() {
	tourDirector := NewTourDirector(NewDefaultTourPlanBuilder())
	tourDirector.CancunTrip()
	tourDirector.LongBeachTrip()
}
