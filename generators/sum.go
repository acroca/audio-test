package generators

// Sum Represents a sum wave generator
type Sum struct {
	a Base
	b Base
}

// ProcessAudio processes the audio
func (sum *Sum) ProcessAudio(out [][2]float32) {
	resA := make([][2]float32, len(out))
	resB := make([][2]float32, len(out))

	sum.a.ProcessAudio(resA)
	sum.b.ProcessAudio(resB)

	for i := range out {
		for j := range out[i] {
			out[i][j] = resA[i][j] + resB[i][j]
			if out[i][j] > 1 {
				out[i][j] = 1
			}
			if out[i][j] < -1 {
				out[i][j] = -1
			}
		}
	}
}

// NewSum returns a new Sum generator
func NewSum(generators ...Base) Base {
	if len(generators) == 0 {
		return NewConst(0)
	}
	if len(generators) == 1 {
		return generators[0]
	}
	acc := &Sum{generators[0], generators[1]}
	for _, gen := range generators[2:] {
		acc = &Sum{acc, gen}
	}
	return acc
}
