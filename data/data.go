package data

import (
	"context"
	"log"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

//go:generate go run generate.go

func PlaySpinLoop(ctx context.Context) {
	file, err := assets.Open("spin-loop.mp3")
	if err != nil || file == nil {
		return
	}
	// Not closing file in favor of closing the streamer, per beep documentation
	streamer, format, err := mp3.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	// log.Printf("%d\n", streamer.Len()) // 801792
	for {
		ended := make(chan bool)
		speaker.Play(beep.Seq(streamer, beep.Callback(func() {
			ended <- true
		})))
		<-ended
		if err := streamer.Seek(0); err != nil {
			log.Fatalf("an error occurred while attempting to loop audio: %s", err)
		}
		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}
