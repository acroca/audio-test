package generators

import "math"

// Triangle Represents a triangle wave generator
type Triangle struct {
	stepL, phaseL float64
	stepR, phaseR float64
}

// ProcessAudio processes the audio
func (triangle *Triangle) ProcessAudio(out [][]float32) {
	for i := range out[0] {
		if triangle.phaseL < 0.25 {
			out[0][i] = float32(triangle.phaseL * 4)
		} else if triangle.phaseL < 0.75 {
			out[0][i] = float32(1 - (4 * (triangle.phaseL - 0.25)))
		} else {
			out[0][i] = float32(-1 + (4 * (triangle.phaseL - 0.75)))
		}
		_, triangle.phaseL = math.Modf(triangle.phaseL + triangle.stepL)

		if triangle.phaseR < 0.25 {
			out[1][i] = float32(triangle.phaseR * 4)
		} else if triangle.phaseR < 0.75 {
			out[1][i] = float32(1 - (4 * (triangle.phaseR - 0.25)))
		} else {
			out[1][i] = float32(-1 + (4 * (triangle.phaseR - 0.75)))
		}
		_, triangle.phaseR = math.Modf(triangle.phaseR + triangle.stepR)
	}
}

// NewTriangle returns a new Triangle generator
func NewTriangle(freqL, freqR, sampleRate float64) *Triangle {
	return &Triangle{freqL / sampleRate, 0, freqR / sampleRate, 0}
}
