package main

import "fmt"

type GameServiceInterface interface {
	StartGame()
}

type GameService struct {
}

func NewGameService() *GameService {
	return &GameService{}
}

func (gs *GameService) StartGame() {
	fmt.Println("게임이 시작되었습니다!")
}
