package main

import (
	"os"
	"time"
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

	win := Window{origin: &Vector{X: 0, Y: 0}, width: 40, height: 20}
	r := NewRenderer(win)
	r.Render()

	g := StartGame()
	for _, p := range g.Positions() {
		r.Put(p, 'o')
	}

	for g.Running {
		input := 0
		if inputReady(fd, 80*time.Millisecond) {
			input = readKey()
		}

		prev := g.Positions()
		g.process(input)
		next := g.Positions()

		for _, p := range prev {
			r.Put(p, ' ')
		}
		for _, p := range next {
			r.Put(p, 'o')
		}
	}
}
