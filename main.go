package main

import (
	"github.com/go-vgo/robotgo"
)

func main() {
	width, height := robotgo.GetScreenSize()
	originalX, originalY := robotgo.GetMousePos()
	robotgo.MoveMouseSmooth(width/2, height/2)
	robotgo.MoveMouse(originalX, originalY)
}
