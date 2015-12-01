package main

import (
	"time"

	"github.com/acroca/audio-test/generators"
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
	time.Sleep(2 * time.Second)
	chk(s.Stop())
}

type stereo struct {
	*portaudio.Stream

	channels []audioGen
}

func newStereo() *stereo {
	s := &stereo{}
	s.channels = []audioGen{
		generators.NewSum(
			// A, C#, and E
			generators.NewMul(
				generators.NewConst(1.0/3),
				generators.NewSine(440, 440, sampleRate),
			),
			generators.NewMul(
				generators.NewConst(1.0/3),
				generators.NewSine(550, 550, sampleRate),
			),
			generators.NewMul(
				generators.NewConst(1.0/3),
				generators.NewSine(660, 660, sampleRate),
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
