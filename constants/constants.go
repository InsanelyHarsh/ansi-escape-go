// Package constants holds values shared across packages: the terminal
// client, the game engine, and (later) the network layer.
package constants

import "time"

// Esc is the ANSI escape sequence prefix.
const Esc = "\033["

// Input key constants describing abstract game input, independent of
// how that input arrives (local terminal keys today, network messages
// later).
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

// Board dimensions.
const (
	BoardWidth  = 40
	BoardHeight = 20
)

// Gameplay tuning.
const (
	PaddleHeight = 3
	MaxBallSpeed = 2
)

// TickInterval is the fixed duration of one game step.
const TickInterval = 150 * time.Millisecond
