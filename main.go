package main

import (
	"math"
	"time"

	g "github.com/acroca/audio-test/generators"
	"github.com/timshannon/go-openal/openal"
)

const (
	sampleRate = 44100
	// maxInt16   = int16((1 << 15) - 1)
	maxInt16 = 32760
)

func main() {
	d := openal.OpenDevice(openal.GetString(openal.DefaultDeviceSpecifier))
	ctx := d.CreateContext()
	ctx.Activate()

	s := openal.NewSource()
	stereo := newStereo()
	var arr = make([][2]int16, sampleRate)
	for i := 0; i < 3; i++ {
		b := openal.NewBuffer()
		stereo.processAudio(arr)
		b.SetDataStereo16(arr, sampleRate)
		s.QueueBuffer(b)
	}

	s.Play()
	time.Sleep(5 * time.Second)
}

// func main2() {
// 	portaudio.Initialize()
// 	defer portaudio.Terminate()
// 	s := newStereo()
// 	s.play()
// 	defer s.Close()
// 	chk(s.Start())
// 	time.Sleep(5 * time.Second)
// 	chk(s.Stop())
// }

type stereo struct {
	channel g.Base
}

func newStereo() *stereo {
	s := &stereo{}
	// s.channel = g.NewSine(440, 440, sampleRate)
	bpm := 130
	samplesPerBeat := int(float32(sampleRate) / float32(float32(bpm)/60.0))

	s.channel = g.NewSum(
		g.NewMul(
			g.NewStop(1*samplesPerBeat),
			g.NewSine(440, 440, sampleRate),
		),
		g.NewMul(
			g.NewPause(1*samplesPerBeat),
			g.NewStop(2*samplesPerBeat),
			g.NewSum(
				g.NewMul(
					g.NewConst(1.0/2),
					g.NewSine(440, 440, sampleRate),
				),
				g.NewMul(
					g.NewConst(1.0/2),
					g.NewSine(550, 550, sampleRate),
				),
			),
		),
		g.NewMul(
			g.NewPause(2*samplesPerBeat),
			g.NewStop(3*samplesPerBeat),
			g.NewSum(
				g.NewMul(
					g.NewConst(1.0/3),
					g.NewSine(440, 440, sampleRate),
				),
				g.NewMul(
					g.NewConst(1.0/3),
					g.NewSine(550, 550, sampleRate),
				),
				g.NewMul(
					g.NewConst(1.0/3),
					g.NewSine(660, 660, sampleRate),
				),
			),
		),
	)
	return s
}

func (g *stereo) processAudio(out [][2]int16) {
	floatOut := make([][2]float32, len(out))
	g.channel.ProcessAudio(floatOut)
	for i := range floatOut {
		out[i][0] = int16(math.Floor(float64(floatOut[i][0] * maxInt16)))
		out[i][1] = int16(math.Floor(float64(floatOut[i][1] * maxInt16)))
	}
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
