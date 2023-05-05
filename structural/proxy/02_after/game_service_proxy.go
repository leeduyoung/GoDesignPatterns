package main

import (
	"fmt"
	"time"
)

type GameServiceProxy struct {
	gameService *GameService
}

func NewGameServiceProxy() *GameServiceProxy {
	return &GameServiceProxy{
		gameService: nil,
	}
}

func (gsp *GameServiceProxy) StartGame() {
	// 다양한 미들웨어 처리
	currentTime := time.Now()
	time.Sleep(time.Second)

	if gsp.gameService == nil {
		gsp.gameService = NewGameService() // LAZY LOADING
	}

	gsp.gameService.StartGame()

	// 게임시작까지 시간 측정
	fmt.Println(time.Since(currentTime))
}
