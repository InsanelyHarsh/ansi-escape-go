package main

import (
	"os"
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

	pos := Vector{X: 1, Y: 1}
	r.Put(pos, '@')

	for {
		key := readKey()

		next := pos
		switch key {
		case KeyUp:
			next.Y--
		case KeyDown:
			next.Y++
		case KeyRight:
			next.X++
		case KeyLeft:
			next.X--
		case KeyEscape:
			return
		default:
			continue
		}

		if !r.Contains(next) {
			continue
		}

		r.Put(pos, ' ')
		r.Put(next, '@')
		pos = next
	}
}
