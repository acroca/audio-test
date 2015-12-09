package generators

// Stop Represents a stop
type Stop struct {
	steps int
}

// ProcessAudio processes the audio
func (stop *Stop) ProcessAudio(out [][]float32) {
	var o float32
	for i := range out[0] {
		if stop.steps > 0 {
			o = 1
			stop.steps--
		} else {
			o = 0
		}
		out[0][i] = o
		out[1][i] = o
	}
}

// NewStop returns a new Stop generator
func NewStop(steps int) *Stop {
	return &Stop{steps}
}
