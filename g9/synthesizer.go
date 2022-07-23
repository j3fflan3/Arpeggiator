package g9

type synthesizer interface {
	Synthesize(frequency float64, duration float64) []float64
}
