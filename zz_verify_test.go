package main

import (
	"fmt"
	"testing"
)

func TestRendererAnchoredMove(t *testing.T) {
	win := Window{origin: &Vector{X: 0, Y: 0}, width: 4, height: 3}
	r := NewRenderer(win)

	fmt.Printf("corners: TL=%+v TR=%+v BL=%+v BR=%+v\n", r.topLeft, r.topRight, r.bottomLeft, r.bottomRight)

	fmt.Println("--- Put in-bounds (2,1) ---")
	r.Put(Vector{X: 2, Y: 1}, '@')

	fmt.Println("--- Put out-of-bounds (10,10), should be dropped (no escape codes/char) ---")
	r.Put(Vector{X: 10, Y: 10}, 'X')

	fmt.Println("--- Draw border ---")
	r.DrawWindow()
}
