package generators

// Adsr Represents a adsr wave manipultor
type Adsr struct {
	state    int
	val      float64
	a, aStep float64
	d, dStep float64
	s        float64
	r, rStep float64
}

// ProcessAudio processes the audio
func (adsr *Adsr) ProcessAudio(out [][]float32) {
	for i := range out[0] {
		out[0][i] = float32(adsr.val)
		out[1][i] = float32(adsr.val)
		switch adsr.state {
		case 0: // attack
			adsr.a--
			if adsr.a <= 0 {
				adsr.state++
			}
		case 1: // decay
			adsr.d--
			if adsr.d <= 0 {
				adsr.state++
			}
		case 2: // sustain
			adsr.s--
			if adsr.s <= 0 {
				adsr.state++
			}
		case 3: // release
			adsr.r--
			if adsr.r <= 0 {
				adsr.state++
			}
		}
		switch adsr.state {
		case 0: // attack
			adsr.val += adsr.aStep
		case 1: // decay
			adsr.val -= adsr.dStep
		case 3: // release
			adsr.val -= adsr.rStep
		}
	}
}

// NewAdsr returns a new Adsr generator
func NewAdsr(a, d, s, sTar, r, samplerate float64) *Adsr {
	return &Adsr{0, 0,
		a * samplerate, 1 / (a * samplerate),
		d * samplerate, (1 - sTar) / (d * samplerate),
		s * samplerate,
		r * samplerate, sTar / (r * samplerate)}
}
