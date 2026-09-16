package main

import "fmt"

type Vector struct {
	X, Y int
}

type Item struct {
	position Vector
	ch       rune
}

type Game struct {
	Width, Height int

	ball        Ball
	leftPaddle  Paddle
	rightPaddle Paddle
	items       []*Item

	LeftScore, RightScore int

	Running bool
}

func (g *Game) ScoreString() string {
	return fmt.Sprintf("%d : %d", g.LeftScore, g.RightScore)
}

func StartGame() *Game {
	g := &Game{
		Running: true,
		Width:   40,
		Height:  20,
	}

	g.ball = Ball{
		Item: Item{
			position: Vector{
				X: g.Width / 2,
				Y: g.Height / 2,
			},
			ch: 'o',
		},
		velocity: randomVelocity(),
		Speed:    1,
	}

	const paddleHeight = 3
	midY := g.Height/2 - paddleHeight/2

	g.leftPaddle = Paddle{
		Item: Item{
			position: Vector{X: 1, Y: midY},
			ch:       '|',
		},
		Width:  1,
		Height: paddleHeight,
		Speed:  1,
	}

	g.rightPaddle = Paddle{
		Item: Item{
			position: Vector{X: g.Width - 2, Y: midY},
			ch:       '|',
		},
		Width:  1,
		Height: paddleHeight,
		Speed:  1,
	}

	return g
}

// AddItem registers a new drawable item in the game's window.
func (g *Game) AddItem(item *Item) {
	g.items = append(g.items, item)
}

// Items returns the current drawable items, for the caller to render.
func (g *Game) Items() []Item {
	gameItems := make([]Item, 0, len(g.items)+1+g.leftPaddle.Height+g.rightPaddle.Height)

	gameItems = append(gameItems, g.ball.Item)
	gameItems = append(gameItems, g.leftPaddle.Items()...)
	gameItems = append(gameItems, g.rightPaddle.Items()...)

	for _, it := range g.items {
		gameItems = append(gameItems, *it)
	}
	return gameItems
}

// process advances the game by one step given this frame's input.
func (g *Game) process(input int) {
	if input == KeyEscape {
		g.Running = false
		return
	}

	switch input {
	case int('w'), int('W'):
		g.movePaddle(&g.leftPaddle, -int(g.leftPaddle.Speed))
	case int('s'), int('S'):
		g.movePaddle(&g.leftPaddle, int(g.leftPaddle.Speed))
	case KeyUp:
		g.movePaddle(&g.rightPaddle, -int(g.rightPaddle.Speed))
	case KeyDown:
		g.movePaddle(&g.rightPaddle, int(g.rightPaddle.Speed))
	}

	g.moveBall()
}

// movePaddle shifts p vertically by dy, clamping so it stays fully
// inside the box.
func (g *Game) movePaddle(p *Paddle, dy int) {
	next := p.position.Y + dy

	if next < 1 || next+p.Height-1 > g.Height-2 {
		return
	}

	p.position.Y = next
}

func (g *Game) moveBall() {
	next := g.ball.GetNextPosition()

	if !g.Contains(Vector{X: next.X, Y: g.ball.position.Y}) {
		var paddle *Paddle
		var defenderScore, attackerScore *int

		if next.X <= 0 {
			paddle, defenderScore, attackerScore = &g.leftPaddle, &g.LeftScore, &g.RightScore
		} else {
			paddle, defenderScore, attackerScore = &g.rightPaddle, &g.RightScore, &g.LeftScore
		}

		if !paddle.overlaps(g.ball.position.Y) {
			*attackerScore++
			g.resetBall()
			return
		}

		*defenderScore++
		g.ball.velocity.X = -g.ball.velocity.X
		g.ball.velocity.Y = clampSpeed(paddle.hitOffset(g.ball.position.Y))
		next.X = g.ball.position.X + g.ball.velocity.X
		next.Y = g.ball.position.Y + g.ball.velocity.Y
	}

	if !g.Contains(Vector{X: g.ball.position.X, Y: next.Y}) {
		g.ball.velocity.Y = -g.ball.velocity.Y
		next.Y = g.ball.position.Y + g.ball.velocity.Y
	}

	g.ball.position = next
}

func (g *Game) resetBall() {
	g.ball.position = Vector{X: g.Width / 2, Y: g.Height / 2}
	g.ball.velocity = randomVelocity()
}

func (g *Game) Contains(pos Vector) bool {
	return pos.X > 0 && pos.X < g.Width-1 &&
		pos.Y > 0 && pos.Y < g.Height-1
}
