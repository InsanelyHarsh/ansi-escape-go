package main

type Paddle struct {
	Item
	Width  int //will be 1
	Height int
	Speed  float64
}

// Items returns one drawable Item per row of the paddle's height.
func (p *Paddle) Items() []Item {
	items := make([]Item, p.Height)
	for i := 0; i < p.Height; i++ {
		items[i] = Item{
			position: Vector{X: p.position.X, Y: p.position.Y + i},
			ch:       p.ch,
		}
	}
	return items
}
