package generators

// Mul Represents a mul wave generator
type Mul struct {
	a Base
	b Base
}

// ProcessAudio processes the audio
func (mul *Mul) ProcessAudio(out [][2]float32) {
	resA := make([][2]float32, len(out))
	resB := make([][2]float32, len(out))

	mul.a.ProcessAudio(resA)
	mul.b.ProcessAudio(resB)

	for i := range out {
		for j := range out[i] {
			out[i][j] = resA[i][j] * resB[i][j]
			if out[i][j] > 1 {
				out[i][j] = 1
			}
			if out[i][j] < -1 {
				out[i][j] = -1
			}
		}
	}
}

// NewMul returns a new Mul generator
func NewMul(generators ...Base) Base {
	if len(generators) == 0 {
		return NewConst(0)
	}
	if len(generators) == 1 {
		return generators[0]
	}
	acc := &Mul{generators[0], generators[1]}
	for _, gen := range generators[2:] {
		acc = &Mul{acc, gen}
	}
	return acc
}
