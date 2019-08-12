package main

import (
	"math"

	"github.com/go-vgo/robotgo"
	"gonum.org/v1/gonum/mat"
)

func main() {
	midScreenMatrix := createMidScreenMatrix()
	originalX, originalY := robotgo.GetMousePos()
	robotgo.MoveMouseSmooth(int(midScreenMatrix.At(0, 0)), int(midScreenMatrix.At(0, 1)))
	robotgo.MoveMouse(originalX, originalY)
}

func createMidScreenMatrix() mat.Matrix {
	screenSizeMatrix := createScreenSizeMatrix()
	midScreenMatrix := mat.NewDense(1, 2, nil)
	midScreenMatrix.Scale(0.5, screenSizeMatrix)
	return midScreenMatrix
}

func createScreenSizeMatrix() mat.Matrix {
	width, height := robotgo.GetScreenSize()
	values := []float64{float64(width), float64(height)}
	return mat.NewDense(1, 2, values)
}

func createRotationMatrix(theta float64) mat.Matrix {
	// Row major order of the rotation matrix
	// https://en.wikipedia.org/wiki/Rotation_matrix
	values := []float64{math.Cos(theta), math.Sin(theta), -math.Sin(theta), math.Cos(theta)}
	return mat.NewDense(2, 2, values)
}
