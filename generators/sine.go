package generators

import "math"

// Sine Represents a sine wave generator
type Sine struct {
	stepL, phaseL float64
	stepR, phaseR float64
}

// ProcessAudio processes the audio
func (sine *Sine) ProcessAudio(out [][2]float32) {
	for i := range out {
		out[i][0] = float32(math.Sin(2 * math.Pi * sine.phaseL))
		_, sine.phaseL = math.Modf(sine.phaseL + sine.stepL)
		out[i][1] = float32(math.Sin(2 * math.Pi * sine.phaseR))
		_, sine.phaseR = math.Modf(sine.phaseR + sine.stepR)
	}
}

// NewSine returns a new Sine generator
func NewSine(freqL, freqR, sampleRate float64) *Sine {
	return &Sine{freqL / sampleRate, 0, freqR / sampleRate, 0}
}
