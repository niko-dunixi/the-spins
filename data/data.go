package data

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

//go:generate go run generate.go
var (
	seekPointBuildOverride = ""
	seekPoint              = 1581000
)

func init() {
	if seekPointBuildOverride == "" {
		return
	}
	if value, err := strconv.Atoi(seekPointBuildOverride); err == nil {
		log.Printf("Overriding seek value from '%d' to '%d'\n", seekPoint, value)
		seekPoint = value
	} else {
		log.Fatalf("Bad seek override value: '%s'", seekPointBuildOverride)
	}
}

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
		if err := streamer.Seek(seekPoint); err != nil {
			log.Fatalf("an error occurred while attempting to loop audio: %s", err)
		}
		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}
