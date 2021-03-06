package generators

import "math"

// Sawtooth Represents a sawtooth wave generator
type Sawtooth struct {
	stepL, phaseL float64
	stepR, phaseR float64
}

// ProcessAudio processes the audio
func (sawtooth *Sawtooth) ProcessAudio(out [][2]float32) {
	for i := range out {
		if sawtooth.phaseL < 0.5 {
			out[i][0] = float32(sawtooth.phaseL * 2)
		} else {
			out[i][0] = float32(-1 + (2 * (sawtooth.phaseL - 0.5)))
		}
		_, sawtooth.phaseL = math.Modf(sawtooth.phaseL + sawtooth.stepL)

		if sawtooth.phaseR < 0.5 {
			out[i][1] = float32(sawtooth.phaseR * 2)
		} else {
			out[i][1] = float32(-1 + (2 * (sawtooth.phaseR - 0.5)))
		}
		_, sawtooth.phaseR = math.Modf(sawtooth.phaseR + sawtooth.stepR)
	}
}

// NewSawtooth returns a new Sawtooth generator
func NewSawtooth(freqL, freqR, sampleRate float64) *Sawtooth {
	return &Sawtooth{freqL / sampleRate, 0, freqR / sampleRate, 0}
}
