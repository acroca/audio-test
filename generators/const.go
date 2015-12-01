package generators

// Const is a generator that sets the output always a fixed value
type Const struct {
	val float32
}

// ProcessAudio processes the audio
func (c *Const) ProcessAudio(out [][]float32) {
	for i := range out[0] {
		out[0][i] = c.val
		out[1][i] = c.val
	}
}

// NewConst returns a new Const generator
func NewConst(val float32) *Const {
	return &Const{val}
}
