package generators

// Const is a generator that sets the output always a fixed value
type Const struct {
	val float32
}

// ProcessAudio processes the audio
func (c *Const) ProcessAudio(out [][2]float32) {
	for i := range out {
		out[i][0] = c.val
		out[i][1] = c.val
	}
}

// NewConst returns a new Const generator
func NewConst(val float32) *Const {
	return &Const{val}
}
