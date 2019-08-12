package main

import (
	"math"
	"time"

	"github.com/go-vgo/robotgo"
	"gonum.org/v1/gonum/mat"
)

func main() {
	midScreenMatrix := createMidScreenMatrix()
	originalX, originalY := robotgo.GetMousePos()
	robotgo.MoveMouseSmooth(int(midScreenMatrix.At(0, 0)), int(midScreenMatrix.At(0, 1)))
	point := mat.NewDense(1, 2, []float64{150, 150})
	theta := 0.0
	for {
		rotationMatrix := createRotationMatrix(theta)
		mouseMatrix := determineMouseMatrix(point, rotationMatrix, midScreenMatrix)
		robotgo.MoveMouse(int(mouseMatrix.At(0, 0)), int(mouseMatrix.At(0, 1)))

		theta += 0.07
		time.Sleep(3 * time.Millisecond)
	}
	robotgo.MoveMouse(originalX, originalY)
}

func determineMouseMatrix(point, rotationMatrix, offset mat.Matrix) mat.Matrix {
	result := mat.NewDense(1, 2, nil)
	result.Mul(point, rotationMatrix)
	result.Add(result, offset)
	return result
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
