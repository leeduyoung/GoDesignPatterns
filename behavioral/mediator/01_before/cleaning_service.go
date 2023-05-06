package main

import "fmt"

type CleaningService struct {
}

func NewCleaningService() *CleaningService {
	return &CleaningService{}
}

func (s CleaningService) GetTower(g *Guest, number int) {
	fmt.Println("[CleaningService] 타워를 준다", number)
}

func (s CleaningService) Clean(r Restaurant) {
	fmt.Println("[CleaningService] 청소를 한다")
}
