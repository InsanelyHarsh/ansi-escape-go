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
	for _, it := range g.Items() {
		r.Put(it.position, it.ch)
	}

	for g.Running {
		input := 0
		if inputReady(fd, 80*time.Millisecond) {
			input = readKey()
		}

		prev := g.Items()
		g.process(input)
		next := g.Items()

		for _, it := range prev {
			r.Put(it.position, ' ')
		}
		for _, it := range next {
			r.Put(it.position, it.ch)
		}
	}
}
