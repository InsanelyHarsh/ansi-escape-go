package main

type Ball struct {
	position Vector
	velocity Vector
}

func (b *Ball) GetNextPosition() Vector {
	return Vector{
		X: b.position.X + b.velocity.X,
		Y: b.position.Y + b.velocity.Y,
	}
}
