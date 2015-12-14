package generators

// Base is the interface implemented by all generators
type Base interface {
	ProcessAudio(out [][2]float32)
}
