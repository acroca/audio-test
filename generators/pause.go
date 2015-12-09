package generators

// Pause Represents a pause
type Pause struct {
	steps int
}

// ProcessAudio processes the audio
func (pause *Pause) ProcessAudio(out [][]float32) {
	var o float32
	for i := range out[0] {
		if pause.steps > 0 {
			o = 0
			pause.steps--
		} else {
			o = 1
		}
		out[0][i] = o
		out[1][i] = o
	}
}

// NewPause returns a new Pause generator
func NewPause(steps int) *Pause {
	return &Pause{steps}
}
