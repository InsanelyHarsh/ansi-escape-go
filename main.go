package main

import (
	"fmt"
	"os"
	"time"
)

func EnterAlternateScreen() {
	fmt.Print("\033[?1049h") //alternate screen
	fmt.Print("\033[2J")     //clear
	fmt.Print("\033[H")      //home
}

func ExitAlternateScreen() {
	fmt.Print("\033[?1049l")
}

const (
	KeyCharacter = 1 << iota
	KeyEscape
	KeyUp
	KeyDown
	KeyLeft
	KeyRight
	KeyEnter
	KeySpace
	KeyBackspace
)

func readKey() int {
	buf := make([]byte, 1)

	_, err := os.Stdin.Read(buf)
	if err != nil {
		return KeyEscape
	}

	switch buf[0] {
	case '\r', '\n':
		return KeyEnter

	case ' ':
		return KeySpace

	case 127:
		return KeyBackspace

	case 27:
		// ESC received. If no follow-up byte shows up quickly, this is a
		// standalone Escape key press, not the start of an escape sequence.
		if !inputReady(int(os.Stdin.Fd()), 25*time.Millisecond) {
			return KeyEscape
		}

		_, err := os.Stdin.Read(buf)
		if err != nil {
			return KeyEscape
		}

		if buf[0] != '[' {
			return KeyEscape
		}

		_, err = os.Stdin.Read(buf)
		if err != nil {
			return KeyEscape
		}

		switch buf[0] {
		case 'A':
			return KeyUp
		case 'B':
			return KeyDown
		case 'C':
			return KeyRight
		case 'D':
			return KeyLeft
		}

		return KeyEscape
	case 'h', 'H':
		return KeyLeft
	case 'j', 'J':
		return KeyDown
	case 'k', 'K':
		return KeyUp
	case 'l', 'L':
		return KeyRight
	default:
		// Printable character.
		return int(buf[0])
	}
}

const (
	// esc = "\u001b["
	esc = "\033["
)

func CursorUp(n int) {
	fmt.Printf("%s%dA", esc, n)
}

func CursorDown(n int) {
	fmt.Printf("%s%dB", esc, n)
}

func CursorLeft(n int) {
	fmt.Printf("%s%dD", esc, n)
}

func CursorRight(n int) {
	fmt.Printf("%s%dC", esc, n)
}

func Delete(n int) {
	fmt.Printf("%s%dP", esc, n)
}

func CursorHome() {
	fmt.Print("\033[H")
}

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
		// case KeyUp, KeyDown, KeyRight, KeyLeft:
		case KeyUp:
			CursorUp(1)
		case KeyDown:
			CursorDown(1)
		case KeyRight:
			CursorRight(1)
		case KeyLeft:
			CursorRight(1)
		case KeyEscape:
			return
		}
	}
}
