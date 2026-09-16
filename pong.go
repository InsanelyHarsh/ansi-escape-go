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

type Paddle struct {
	X, Y int
	// int height;
}

type Ball struct {
	position Vector
	velocity Vector
}

type Game struct {
	Width, Height int

	// Paddle left
	// Paddle right
	ball Ball

	Running bool
}

// take input as process game state
func (g *Game) process(input int) {
	for true {
		if !g.Running {

		}

	}

}
