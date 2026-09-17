package main

import (
	"fmt"
	"os"
	"time"

	"github.com/insanelyharsh/ansi-escape-go/constants"
)

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
		return constants.KeyEscape
	}

	switch buf[0] {
	case '\r', '\n':
		return constants.KeyEnter

	case ' ':
		return constants.KeySpace

	case 127:
		return constants.KeyBackspace

	case 27:
		// ESC received. If no follow-up byte shows up quickly, this is a
		// standalone Escape key press, not the start of an escape sequence.
		if !inputReady(int(os.Stdin.Fd()), 25*time.Millisecond) {
			return constants.KeyEscape
		}

		_, err := os.Stdin.Read(buf)
		if err != nil {
			return constants.KeyEscape
		}

		if buf[0] != '[' {
			return constants.KeyEscape
		}

		_, err = os.Stdin.Read(buf)
		if err != nil {
			return constants.KeyEscape
		}

		switch buf[0] {
		case 'A':
			return constants.KeyUp
		case 'B':
			return constants.KeyDown
		case 'C':
			return constants.KeyRight
		case 'D':
			return constants.KeyLeft
		}

		return constants.KeyEscape
	case 'h', 'H':
		return constants.KeyLeft
	case 'j', 'J':
		return constants.KeyDown
	case 'k', 'K':
		return constants.KeyUp
	case 'l', 'L':
		return constants.KeyRight
	default:
		// Printable character.
		return int(buf[0])
	}
}

func CursorUp(n int) {
	fmt.Printf("%s%dA", constants.Esc, n)
}

func CursorDown(n int) {
	fmt.Printf("%s%dB", constants.Esc, n)
}

func CursorLeft(n int) {
	fmt.Printf("%s%dD", constants.Esc, n)
}

func CursorRight(n int) {
	fmt.Printf("%s%dC", constants.Esc, n)
}

func Delete(n int) {
	fmt.Printf("%s%dP", constants.Esc, n)
}

func CursorHome() {
	fmt.Print("\033[H")
}
