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
