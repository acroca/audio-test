package main

import (
	"time"

	g "github.com/acroca/audio-test/generators"
	"github.com/gordonklaus/portaudio"
)

const sampleRate = 44100

type audioGen interface {
	ProcessAudio(out [][]float32)
}

func main() {
	portaudio.Initialize()
	defer portaudio.Terminate()
	s := newStereo()
	defer s.Close()
	chk(s.Start())
	time.Sleep(5 * time.Second)
	chk(s.Stop())
}

type stereo struct {
	*portaudio.Stream

	channels []audioGen
}

func newStereo() *stereo {
	s := &stereo{}
	bpm := 130
	samplesPerBeat := int(float32(sampleRate) / float32(float32(bpm)/60.0))

	s.channels = []audioGen{
		g.NewSum(
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
		),
	}
	var err error
	s.Stream, err = portaudio.OpenDefaultStream(0, 2, sampleRate, 0, s.processAudio)
	chk(err)
	return s
}

func (g *stereo) processAudio(out [][]float32) {
	for _, channel := range g.channels {
		channel.ProcessAudio(out)
	}
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
