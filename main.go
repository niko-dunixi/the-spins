package main

import (
	// "context"
	// "math"
	// "os"
	// "os/signal"
	// "time"

	"github.com/faiface/pixel"
	"github.com/go-vgo/robotgo"
	"golang.org/x/image/colornames"

	// "github.com/paul-nelson-baker/the-spins/data"
	// "gonum.org/v1/gonum/mat"

	"github.com/faiface/pixel/pixelgl"
)

func main() {
	// Capture the desktop as a screenCaptureBitmap
	width, height := robotgo.GetScreenSize()
	screenCaptureBitmap := robotgo.CaptureScreen()
	defer robotgo.FreeBitmap(screenCaptureBitmap)
	// Create a window and display said screenCaptureBitmap
	pixelgl.Run(func() {
		cfg := pixelgl.WindowConfig{
			Undecorated: true,
			//AlwaysOnTop: true,
			VSync:       true,
			Bounds: pixel.Rect{
				Max: pixel.Vec{
					X: float64(width),
					Y: float64(height),
				},
			},
			Monitor: pixelgl.PrimaryMonitor(),
		}
		win, err := pixelgl.NewWindow(cfg)
		if err != nil {
			panic(err)
		}
		win.Clear(colornames.Black)
		for !win.Closed() {
			win.Update()
		}
	})
}

// func main() {

// 	c1, cancel := context.WithCancel(context.Background())
// 	exitCh := make(chan struct{})

// 	originalX, originalY := robotgo.GetMousePos()

// 	go func(ctx context.Context) {
// 		midScreenMatrix := createMidScreenMatrix()
// 		go data.PlaySpinLoop(ctx)
// 		robotgo.MoveMouseSmooth(int(midScreenMatrix.At(0, 0)), int(midScreenMatrix.At(0, 1)))
// 		time.Sleep(600 * time.Millisecond)
// 		point := mat.NewDense(1, 2, []float64{75, 75})
// 		theta := 0.0
// 		for {
// 			// Mouse stuff
// 			mouseMatrix := determineMouseMatrix(point, midScreenMatrix, theta)
// 			robotgo.MoveMouse(int(mouseMatrix.At(0, 0)), int(mouseMatrix.At(0, 1)))
// 			theta += 0.07
// 			time.Sleep(3 * time.Millisecond)
// 			// Loop exit logic
// 			select {
// 			case <-ctx.Done():
// 				time.Sleep(500 * time.Millisecond)
// 				exitCh <- struct{}{}
// 				return
// 			default:
// 			}
// 		}
// 	}(c1)

// 	signalCh := make(chan os.Signal, 1)
// 	signal.Notify(signalCh, os.Interrupt)
// 	go func() {
// 		select {
// 		case <-signalCh:
// 			cancel()
// 			return
// 		}
// 	}()
// 	<-exitCh

// 	robotgo.MoveMouse(originalX, originalY)
// }

// func determineMouseMatrix(point, offset mat.Matrix, theta float64) mat.Matrix {
// 	result := mat.NewDense(1, 2, nil)
// 	rotationMatrix := createRotationMatrix(theta)
// 	result.Mul(point, rotationMatrix)
// 	result.Add(result, offset)
// 	return result
// }

// func createMidScreenMatrix() mat.Matrix {
// 	screenSizeMatrix := createScreenSizeMatrix()
// 	midScreenMatrix := mat.NewDense(1, 2, nil)
// 	midScreenMatrix.Scale(0.5, screenSizeMatrix)
// 	return midScreenMatrix
// }

// func createScreenSizeMatrix() mat.Matrix {
// 	width, height := robotgo.GetScreenSize()
// 	values := []float64{float64(width), float64(height)}
// 	return mat.NewDense(1, 2, values)
// }

// func createRotationMatrix(theta float64) mat.Matrix {
// 	// Row major order of the rotation matrix
// 	// https://en.wikipedia.org/wiki/Rotation_matrix
// 	values := []float64{math.Cos(theta), math.Sin(theta), -math.Sin(theta), math.Cos(theta)}
// 	return mat.NewDense(2, 2, values)
// }
