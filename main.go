package main

import (
	"context"
	"math"
	"os"
	"os/signal"
	"time"

	"github.com/go-vgo/robotgo"
	"gonum.org/v1/gonum/mat"
)

func main() {
	c1, cancel := context.WithCancel(context.Background())
	exitCh := make(chan struct{})

	originalX, originalY := robotgo.GetMousePos()

	go func(ctx context.Context) {
		midScreenMatrix := createMidScreenMatrix()
		robotgo.MoveMouseSmooth(int(midScreenMatrix.At(0, 0)), int(midScreenMatrix.At(0, 1)))
		theta := 0.0
		for {
			// Mouse stuff
			mouseMatrix := determineMouseMatrix(midScreenMatrix, theta)
			robotgo.MoveMouse(int(mouseMatrix.At(0, 0)), int(mouseMatrix.At(0, 1)))
			theta += 0.07
			time.Sleep(3 * time.Millisecond)
			// Loop exit logic
			select {
			case <-ctx.Done():
				time.Sleep(500 * time.Millisecond)
				exitCh <- struct{}{}
				return
			default:
			}
		}
	}(c1)

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)
	go func() {
		select {
		case <-signalCh:
			cancel()
			return
		}
	}()
	<-exitCh

	robotgo.MoveMouse(originalX, originalY)
}

func determineMouseMatrix(offset mat.Matrix, theta float64) mat.Matrix {
	point := mat.NewDense(1, 3, []float64{75, 75, 75})

	result := mat.NewDense(1, 3, nil)
	xAxisRotationMatrix := createXAxisRotationMatrix(theta)
	result.Mul(point, xAxisRotationMatrix)
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

func createXAxisRotationMatrix(theta float64) mat.Matrix {
	// Row major order of the rotation matrix
	// https://en.wikipedia.org/wiki/Rotation_matrix
	// https://en.wikipedia.org/wiki/Rotation_matrix#In_three_dimensions
	// https://math.stackexchange.com/questions/2305792/3d-projection-on-a-2d-plane-weak-maths-ressources
	// Normalize: https://youtu.be/ih20l3pJoeU?list=PLh9hqdYTkzv2ONWRPiMehvG46VWsRBb1O&t=1412
	// https://youtu.be/ih20l3pJoeU?list=PLh9hqdYTkzv2ONWRPiMehvG46VWsRBb1O&t=1679
	rad := asRadians(theta)
	values := []float64{
		1, 0, 0, // column 1
		0, math.Cos(rad), math.Sin(rad), // column 2
		0, -math.Sin(rad), math.Cos(rad), // column 3
	}
	return mat.NewDense(3, 3, values)
}

func createProjectionMatrix(width int, height int, fov float64, zNear float64, zFar float64) mat.Matrix {
	aspectRatio := float64(height) / float64(width)
	return mat.NewDense(4, 4, []float64{
		aspectRatio * math.Tan(asRadians(fov/2)), 0, 0, 0,
		0, 1 / math.Tan(asRadians(fov/2)), 0, 0,
		0, 0, zFar / (zFar - zNear), -(zFar*zNear)/zFar - zNear,
		0, 0, 1, 0,
	})
}

func asRadians(theta float64) (radians float64) {
	return theta / 180.0 * math.Pi
}
