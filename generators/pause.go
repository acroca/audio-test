package generators

// Pause Represents a pause
type Pause struct {
	steps int
}

// ProcessAudio processes the audio
func (pause *Pause) ProcessAudio(out [][2]float32) {
	var o float32
	for i := range out {
		if pause.steps > 0 {
			o = 0
			pause.steps--
		} else {
			o = 1
		}
		out[i][0] = o
		out[i][1] = o
	}
}

// NewPause returns a new Pause generator
func NewPause(steps int) *Pause {
	return &Pause{steps}
}
