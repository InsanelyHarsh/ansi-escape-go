package main

import (
	"fmt"

	"github.com/insanelyharsh/ansi-escape-go/constants"
)

type Renderer struct {
	topLeft     Vector
	topRight    Vector
	bottomLeft  Vector
	bottomRight Vector
}

func NewRenderer(win Window) *Renderer {
	ox, oy := 0, 0
	if win.origin != nil {
		ox, oy = win.origin.X, win.origin.Y
	}

	return &Renderer{
		topLeft:     Vector{X: ox, Y: oy},
		topRight:    Vector{X: ox + win.width - 1, Y: oy},
		bottomLeft:  Vector{X: ox, Y: oy + win.height - 1},
		bottomRight: Vector{X: ox + win.width - 1, Y: oy + win.height - 1},
	}
}

func (r *Renderer) Clear() {
	fmt.Print(constants.Esc + "2J")
	CursorHome()
}

func (r *Renderer) Render() {
	r.Clear()
	r.DrawWindow()
}

func (r *Renderer) DrawWindow() {
	for x := r.topLeft.X; x <= r.topRight.X; x++ {
		r.Put(Vector{X: x, Y: r.topLeft.Y}, '-')
		r.Put(Vector{X: x, Y: r.bottomLeft.Y}, '-')
	}

	for y := r.topLeft.Y; y <= r.bottomLeft.Y; y++ {
		r.Put(Vector{X: r.topLeft.X, Y: y}, '|')
		r.Put(Vector{X: r.topRight.X, Y: y}, '|')
	}

	r.Put(r.topLeft, '+')
	r.Put(r.topRight, '+')
	r.Put(r.bottomLeft, '+')
	r.Put(r.bottomRight, '+')
}

func (r *Renderer) containsWin(pos Vector) bool {
	return pos.X >= r.topLeft.X && pos.X <= r.topRight.X &&
		pos.Y >= r.topLeft.Y && pos.Y <= r.bottomLeft.Y
}

func (r *Renderer) Contains(pos Vector) bool {
	return pos.X > r.topLeft.X && pos.X < r.topRight.X &&
		pos.Y > r.topLeft.Y && pos.Y < r.bottomLeft.Y
}

func (r *Renderer) Put(pos Vector, ch rune) {
	if !r.containsWin(pos) {
		return
	}
	r.Move(pos)
	fmt.Printf("%c", ch)
}

// PutString draws s starting at pos, one character per column.
func (r *Renderer) PutString(pos Vector, s string) {
	for i, ch := range s {
		r.Put(Vector{X: pos.X + i, Y: pos.Y}, ch)
	}
}

func (r *Renderer) Move(pos Vector) {
	CursorHome()

	if pos.Y > 0 {
		CursorDown(pos.Y)
	}
	if pos.X > 0 {
		CursorRight(pos.X)
	}
}
