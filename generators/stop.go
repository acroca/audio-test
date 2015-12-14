package generators

// Stop Represents a stop
type Stop struct {
	steps int
}

// ProcessAudio processes the audio
func (stop *Stop) ProcessAudio(out [][2]float32) {
	var o float32
	for i := range out {
		if stop.steps > 0 {
			o = 1
			stop.steps--
		} else {
			o = 0
		}
		out[i][0] = o
		out[i][1] = o
	}
}

// NewStop returns a new Stop generator
func NewStop(steps int) *Stop {
	return &Stop{steps}
}
