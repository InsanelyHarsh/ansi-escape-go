package main

import "math/rand"

const maxBallSpeed = 2

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

// randomVelocity returns a random velocity capped at maxBallSpeed per
// axis, with a nonzero X so the ball always drifts toward a paddle.
func randomVelocity() Vector {
	x := rand.Intn(maxBallSpeed) + 1
	if rand.Intn(2) == 0 {
		x = -x
	}

	y := rand.Intn(2*maxBallSpeed+1) - maxBallSpeed

	return Vector{X: x, Y: y}
}

func clampSpeed(v int) int {
	if v > maxBallSpeed {
		return maxBallSpeed
	}
	if v < -maxBallSpeed {
		return -maxBallSpeed
	}
	return v
}
