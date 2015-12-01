package generators

import "math"

// Sine Represents a sine wave generator
type Sine struct {
	stepL, phaseL float64
	stepR, phaseR float64
}

// ProcessAudio processes the audio
func (sine *Sine) ProcessAudio(out [][]float32) {
	for i := range out[0] {
		out[0][i] = float32(math.Sin(2 * math.Pi * sine.phaseL))
		_, sine.phaseL = math.Modf(sine.phaseL + sine.stepL)
		out[1][i] = float32(math.Sin(2 * math.Pi * sine.phaseR))
		_, sine.phaseR = math.Modf(sine.phaseR + sine.stepR)
	}
}

// NewSine returns a new Sine generator
func NewSine(freqL, freqR, sampleRate float64) *Sine {
	return &Sine{freqL / sampleRate, 0, freqR / sampleRate, 0}
}
