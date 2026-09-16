package main

import (
	"os"
)

func main() {

	EnterAlternateScreen()
	defer EnterAlternateScreen()

	fd := int(os.Stdin.Fd())

	old, err := enableRawMode(fd)
	if err != nil {
		panic(err)
	}
	defer restore(fd, old)

	// fmt.Println("Press keys. ESC exits.")

	for {
		key := readKey()

		switch key {
		case KeyUp:
			CursorUp(1)
		case KeyDown:
			CursorDown(1)
		case KeyRight:
			CursorRight(1)
		case KeyLeft:
			CursorLeft(1)
		case KeyEscape:
			return
		}
	}
}
