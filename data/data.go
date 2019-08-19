package data

import (
	"log"
	"time"

	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

//go:generate go run generate.go

func PlaySpinLoop() {
	file, err := assets.Open("spin-loop.mp3")
	if err != nil || file == nil {
		return
	}
	streamer, format, err := mp3.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)
	select {}
}
