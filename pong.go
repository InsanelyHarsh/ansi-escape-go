package main

type Vector struct {
	X, Y int
}

type Window struct {
	origin        *Vector
	height, width int
}

// TODO: take config
func NewWindow() *Window {
	return &Window{}
}

type Game struct {
	Width, Height int
	ball          Ball
	Running       bool
}

func StartGame() *Game {
	return &Game{
		Running: true,
		Width:   40,
		Height:  20,
		ball: Ball{
			velocity: Vector{
				X: 1,
				Y: 0,
			},
			position: Vector{
				X: 1,
				Y: 10,
			},
		},
	}
}

// process advances the game by one step given this frame's input.
func (g *Game) process(input int) {
	if input&KeyEscape != 0 {
		g.Running = false
		return
	}

	//TODO: paddle collision handling
	//upon collision
	//velocity is changed

	g.moveBall()
}

func (g *Game) Positions() []Vector {
	return []Vector{g.ball.position}
}

func (g *Game) moveBall() {
	next := g.ball.GetNextPosition()

	if !g.Contains(Vector{X: next.X, Y: g.ball.position.Y}) {
		g.ball.velocity.X = -g.ball.velocity.X
		next.X = g.ball.position.X + g.ball.velocity.X
	}

	if !g.Contains(Vector{X: g.ball.position.X, Y: next.Y}) {
		g.ball.velocity.Y = -g.ball.velocity.Y
		next.Y = g.ball.position.Y + g.ball.velocity.Y
	}

	g.ball.position = next
}

func (g *Game) Contains(pos Vector) bool {
	return pos.X > 0 && pos.X < g.Width-1 &&
		pos.Y > 0 && pos.Y < g.Height-1
}
