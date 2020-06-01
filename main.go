package main

import (
	// "context"
	// "math"
	// "os"
	// "os/signal"
	// "time"

	"bytes"
	"github.com/faiface/pixel"
	"github.com/go-vgo/robotgo"
	"golang.org/x/image/bmp"
	"image"
	"image/png"
	"io/ioutil"
	"strings"
	"time"

	// "github.com/paul-nelson-baker/the-spins/data"
	// "gonum.org/v1/gonum/mat"

	"github.com/faiface/pixel/pixelgl"
)

var fragmentShader = `
#version 330 core

out vec4 fragColor;

uniform sampler2D uTexture;
uniform vec4 uTexBounds;

// custom uniforms
uniform float uSpeed;
uniform float uTime;

void main() {
    vec2 t = gl_FragCoord.xy / uTexBounds.zw;
	vec3 influence = texture(uTexture, t).rgb;

    if (influence.r + influence.g + influence.b > 0.3) {
		t.y += cos(t.x * 40.0 + (uTime * uSpeed))*0.005;
		t.x += cos(t.y * 40.0 + (uTime * uSpeed))*0.01;
	}

    vec3 col = texture(uTexture, t).rgb;
	fragColor = vec4(col * vec3(0.6, 0.6, 1.2),1.0);
}
`

func main() {
	// Capture the screen in the form of image.Image
	screenCaptureImage, err := captureScreenPicture()
	if err != nil {
		panic(err)
	}
	// Save it for debugging purposes
	if err := saveImage(screenCaptureImage, "screen-capture-test"); err != nil {
		panic(err)
	}
	// Convert the image.Image to a pixel.Sprite
	screenCapturePicture := pixel.PictureDataFromImage(screenCaptureImage)
	screenCaptureBounds := imageBounds(screenCaptureImage)
	screenCaptureSprite := pixel.NewSprite(screenCapturePicture, screenCaptureBounds)
	// Create a window and display the desktop which has been captured as a sprite
	var uTime, uSpeed float32
	pixelgl.Run(func() {
		cfg := pixelgl.WindowConfig{
			Undecorated: true,
			//AlwaysOnTop: true,
			VSync:   true,
			Bounds:  screenCaptureBounds,
			Monitor: pixelgl.PrimaryMonitor(),
		}
		win, err := pixelgl.NewWindow(cfg)
		if err != nil {
			panic(err)
		}
		win.Canvas().SetUniform("uTime", &uTime)
		win.Canvas().SetUniform("uSpeed", &uSpeed)
		uSpeed = 5.0
		win.Canvas().SetFragmentShader(fragmentShader)

		start := time.Now()
		for !win.Closed() {
			win.Clear(pixel.RGB(0, 0, 0))
			screenCaptureSprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
			uTime = float32(time.Since(start).Seconds())
			if win.Pressed(pixelgl.KeyRight) {
				uSpeed += 0.1
			}
			if win.Pressed(pixelgl.KeyLeft) {
				uSpeed -= 0.1
			}
			if win.Pressed(pixelgl.KeyEscape) {
				win.SetClosed(true)
			}
			win.Update()
		}
	})
}

func captureScreenPicture() (image.Image, error) {
	screenCaptureBitmap := robotgo.CaptureScreen()
	// This free should be safe, because we're copying the image as we convert it to our target type
	defer robotgo.FreeBitmap(screenCaptureBitmap)
	screenCaptureBytes := robotgo.ToBitmapBytes(screenCaptureBitmap)
	screenCaptureImage, err := bmp.Decode(bytes.NewReader(screenCaptureBytes))
	if err != nil {
		return nil, err
	}
	return screenCaptureImage, nil
}

func saveImage(img image.Image, name string) error {
	byteBuffer := &bytes.Buffer{}
	err := png.Encode(byteBuffer, img)
	if err != nil {
		return err
	}
	if !strings.HasSuffix(name, ".png") {
		name += ".png"
	}
	return ioutil.WriteFile(name, byteBuffer.Bytes(), 0644)
}

func imageBounds(img image.Image) pixel.Rect {
	min := img.Bounds().Min
	max := img.Bounds().Max
	return pixel.R(float64(min.X), float64(min.Y), float64(max.X), float64(max.Y))
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
