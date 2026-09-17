package main

import (
	"math/rand"

	"github.com/insanelyharsh/ansi-escape-go/constants"
)

type Ball struct {
	Item
	velocity Vector
	Speed    float64
}

func (b *Ball) GetNextPosition() Vector {
	return Vector{
		X: b.position.X + b.velocity.X,
		Y: b.position.Y + b.velocity.Y,
	}
}

// randomVelocity returns a random velocity capped at
// constants.MaxBallSpeed per axis, with a nonzero X so the ball
// always drifts toward a paddle.
func randomVelocity() Vector {
	x := rand.Intn(constants.MaxBallSpeed) + 1
	if rand.Intn(2) == 0 {
		x = -x
	}

	y := rand.Intn(2*constants.MaxBallSpeed+1) - constants.MaxBallSpeed

	return Vector{X: x, Y: y}
}

// clampSpeed keeps v within [-constants.MaxBallSpeed, constants.MaxBallSpeed].
func clampSpeed(v int) int {
	if v > constants.MaxBallSpeed {
		return constants.MaxBallSpeed
	}
	if v < -constants.MaxBallSpeed {
		return -constants.MaxBallSpeed
	}
	return v
}
