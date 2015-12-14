package generators

import "math"

// Square Represents a square wave generator
type Square struct {
	stepL, phaseL float64
	stepR, phaseR float64
}

// ProcessAudio processes the audio
func (square *Square) ProcessAudio(out [][2]float32) {
	for i := range out {
		if square.phaseL < 0.5 {
			out[i][0] = 1
		} else {
			out[i][0] = -1
		}
		_, square.phaseL = math.Modf(square.phaseL + square.stepL)
		if square.phaseR < 0.5 {
			out[i][1] = 1
		} else {
			out[i][1] = -1
		}
		_, square.phaseR = math.Modf(square.phaseR + square.stepR)
	}
}

// NewSquare returns a new Square generator
func NewSquare(freqL, freqR, sampleRate float64) *Square {
	return &Square{freqL / sampleRate, 0, freqR / sampleRate, 0}
}
