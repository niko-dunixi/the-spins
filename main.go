package main

import (
	"context"
	"math"
	"os"
	"os/signal"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"gonum.org/v1/gonum/mat"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	exitChannel := make(chan struct{})

	originalX, originalY := robotgo.GetMousePos()

	go func(ctx context.Context) {
		midScreenMatrix := createMidScreenMatrix()
		go PlaySpinLoop(ctx)
		robotgo.MoveMouseSmooth(int(midScreenMatrix.At(0, 0)), int(midScreenMatrix.At(0, 1)))
		time.Sleep(600 * time.Millisecond)
		point := mat.NewDense(1, 2, []float64{75, 75})
		theta := 0.0
		for {
			// Mouse stuff
			mouseMatrix := determineMouseMatrix(point, midScreenMatrix, theta)
			robotgo.MoveMouse(int(mouseMatrix.At(0, 0)), int(mouseMatrix.At(0, 1)))
			theta += 0.07
			time.Sleep(3 * time.Millisecond)
			// Loop exit logic
			select {
			case <-ctx.Done():
				time.Sleep(500 * time.Millisecond)
				exitChannel <- struct{}{}
				return
			default:
			}
		}
	}(ctx)

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)

	go func() {
		// Always trigger the context cancelation because it means
		// an event has killed our loop and we're cascading down the
		// line
		defer cancel()

		// Start reading user input events
		userEventChannel := hook.Start()
		defer hook.End()

		// Debug timeout for times I break the exit logic as I'm
		// experimenting with code.
		// debugTimeout := time.After(time.Second * 60)

		// Wait until the user triggers an exit
		for {
			select {
			case <-signalCh:
				return
			// case <-debugTimeout:
			// return
			case currentEvent := <-userEventChannel:
				if isEscape := currentEvent.Keycode == robotgo.Keycode[`esc`]; isEscape {
					return
				}
			}
		}
	}()
	<-exitChannel

	robotgo.MoveMouse(originalX, originalY)
}

func determineMouseMatrix(point, offset mat.Matrix, theta float64) mat.Matrix {
	result := mat.NewDense(1, 2, nil)
	rotationMatrix := createRotationMatrix(theta)
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
