package g9

import "github.com/faiface/beep"

type Guitar struct {
	sampleRate float64
	synth      synthesizer
}

func NewGuitar(sampleRate int, synth synthesizer) *Guitar {
	return &Guitar{float64(sampleRate), synth}
}

func (g *Guitar) Pluck(note float64, duration float64) beep.Streamer {
	return newSound(g.synth, note, duration)
}

func (g *Guitar) Chord(notes []float64, duration, delay float64) beep.Streamer {
	streamers := make([]beep.Streamer, len(notes))
	for i, note := range notes {
		silence := beep.Silence(int(g.sampleRate * delay * float64(i)))
		sound := newSound(g.synth, note, duration-delay*float64(i))
		streamers[i] = beep.Seq(silence, sound)
	}

	return beep.Mix(streamers...)
}

func (g *Guitar) Silence(duration float64) beep.Streamer {
	return beep.Silence(int(g.sampleRate * duration))
}
