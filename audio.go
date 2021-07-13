package main

import (
	"context"
	"embed"
	"log"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

//go:embed assets/*
var assetContentFS embed.FS

func PlaySpinLoop(_ context.Context) {
	file, err := assetContentFS.Open("assets/spin-loop.mp3")
	if err != nil {
		return
	} else if file == nil {
		return
	}
	// Not closing file in favor of closing the streamer, per beep documentation
	streamer, format, err := mp3.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	songEndedChannel := make(chan struct{})
	defer close(songEndedChannel)
	for {
		speaker.Play(beep.Seq(streamer, beep.Callback(func() {
			songEndedChannel <- struct{}{}
		})))
		<-songEndedChannel
		if err := streamer.Seek(0); err != nil {
			log.Fatalf("an error occurred while attempting to loop audio: %s", err)
		}
	}
}
