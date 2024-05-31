package main

import "fmt"

func main() {
	tourDirector := NewTourDirector(NewTourPlanBuilder())
	cancunTrip := tourDirector.CancunTrip()
	fmt.Println(cancunTrip)

	longBeachTrip := tourDirector.LongBeachTrip()
	fmt.Println(longBeachTrip)
}
