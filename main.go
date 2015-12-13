package main

import (
	"math"
	"time"

	g "github.com/acroca/audio-test/generators"
	"github.com/gordonklaus/portaudio"
	"github.com/timshannon/go-openal/openal"
)

const (
	sampleRate = 44100
	// maxInt16   = int16((1 << 15) - 1)
	maxInt16 = 32760
)

type audioGen interface {
	ProcessAudio(out [][]float32)
}

func main() {
	d := openal.OpenDevice(openal.GetString(openal.DefaultDeviceSpecifier))
	ctx := d.CreateContext()
	ctx.Activate()

	s := openal.NewSource()
	stereo := newStereo()
	var arr = make([][]float32, 2)
	arr[0] = make([]float32, sampleRate)
	arr[1] = make([]float32, sampleRate)
	var arrI = make([][2]int16, sampleRate)
	for i := 0; i < 3; i++ {
		b := openal.NewBuffer()
		stereo.processAudio(arr)
		for i := 0; i < sampleRate; i++ {
			arrI[i][0] = int16(math.Floor(float64(arr[0][i] * float32(maxInt16))))
			arrI[i][1] = int16(math.Floor(float64(arr[1][i] * float32(maxInt16))))
		}
		b.SetDataStereo16(arrI, sampleRate)
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
	*portaudio.Stream

	channel audioGen
}

func newStereo() *stereo {
	s := &stereo{}
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

func (s *stereo) play() {
	var err error
	s.Stream, err = portaudio.OpenDefaultStream(0, 2, sampleRate, 0, s.processAudio)
	chk(err)
}
func (g *stereo) processAudio(out [][]float32) {
	g.channel.ProcessAudio(out)
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
