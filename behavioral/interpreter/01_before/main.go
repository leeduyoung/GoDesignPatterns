package main

import "fmt"

func main() {
	// 현재 게임의 레벨과 점수
	game := new(Game)
	game.SetLevel(5)
	game.SetScore(120)

	// 현재 시점의 레벨과 점수 저장
	currentLevel := game.Level()
	currentScore := game.Score()

	// 다시 저장해둔 시점의 레벨과 점수 불러오기
	newGame := new(Game)
	newGame.SetLevel(currentLevel)
	newGame.SetScore(currentScore)

	fmt.Println(game.Level(), game.Score()) // 5, 120
}
