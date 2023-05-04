package main

type Game struct {
	level int
	score int
}

func (g *Game) Level() int {
	return g.level
}

func (g *Game) SetLevel(level int) {
	g.level = level
}

func (g *Game) Score() int {
	return g.score
}

func (g *Game) SetScore(score int) {
	g.score = score
}

func (g *Game) Save() *GameSave {
	return NewGameSave(g.level, g.score)
}

func (g *Game) ReLoad(gameSave *GameSave) {
	g.level = gameSave.Level()
	g.score = gameSave.Score()
}
