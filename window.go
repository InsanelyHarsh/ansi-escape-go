package main

type Window struct {
	origin        *Vector
	height, width int
}

// TODO: take config
func NewWindow() *Window {
	return &Window{}
}
