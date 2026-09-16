package main

import "fmt"

func EnterAlternateScreen() {
	fmt.Print("\033[?1049h") //alternate screen
	fmt.Print("\033[2J")     //clear
	fmt.Print("\033[H")      //home
}

func ExitAlternateScreen() {
	fmt.Print("\033[?1049l")
}
