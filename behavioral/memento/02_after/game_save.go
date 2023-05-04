package main

type GameSave struct {
	level int
	score int
}

func NewGameSave(level, score int) *GameSave {
	return &GameSave{
		level: level,
		score: score,
	}
}

func (gs GameSave) Level() int {
	return gs.level
}

func (gs GameSave) Score() int {
	return gs.score
}
