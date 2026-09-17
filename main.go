package main

import (
	"os"
	"time"

	"github.com/insanelyharsh/ansi-escape-go/constants"
	// "github.com/insanelyharsh/ansi-escape-go/game"
)

func main() {

	EnterAlternateScreen()
	defer ExitAlternateScreen()

	fd := int(os.Stdin.Fd())

	old, err := enableRawMode(fd)
	if err != nil {
		panic(err)
	}
	defer restore(fd, old)

	// win := Window{origin: &game.Vector{X: 0, Y: 0}, width: constants.BoardWidth, height: constants.BoardHeight}
	win := Window{origin: &Vector{X: 0, Y: 0}, width: constants.BoardWidth, height: constants.BoardHeight}
	r := NewRenderer(win)
	r.Render()

	g := StartGame()
	for _, it := range g.Items() {
		r.Put(it.position, it.ch)
	}
	drawScore(r, g)

	for g.Running {
		tickStart := time.Now()

		input := 0
		if inputReady(fd, 0) {
			input = readKey()
		}

		prev := g.Items()
		g.Process(input)
		next := g.Items()

		for _, it := range prev {
			r.Put(it.position, ' ')
		}
		for _, it := range next {
			r.Put(it.position, it.ch)
		}
		drawScore(r, g)

		if elapsed := time.Since(tickStart); elapsed < constants.TickInterval {
			time.Sleep(constants.TickInterval - elapsed)
		}
	}
}

// drawScore renders the current score in the top-right corner of the
// game's window, overwriting that stretch of the border.
func drawScore(r *Renderer, g *Game) {
	score := g.ScoreString()
	r.PutString(Vector{X: g.Width - 2 - len(score), Y: 0}, score)
}
