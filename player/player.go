package player

import (
	"fmt"
	"os"

	"github.com/faiface/beep"
	"github.com/j3fflan3/arpeggiator/g9"
	"github.com/j3fflan3/arpeggiator/karplusstrong"
	"gopkg.in/yaml.v3"
)

var Streams []beep.Streamer
var Sequence []NoteEvent
var SampleRate int = 44100
var CurrentSong Song

type NoteEvent struct {
	Note     float64 `yaml:"note"`
	Duration float64 `yaml:"duration"`
}

type ChordEvent struct {
	Notes []NoteEvent
	Name  string
}

type LeadSheet struct {
	Note     string
	Octave   int
	Duration string
}

type ChordSheet struct {
	Chord    []ChordNote `yaml:"chord"`
	Duration string      `yaml:"duration"`
}

type ChordNote struct {
	Note   string `yaml:"note"`
	Octave int    `yaml:"octave"`
}

type Song struct {
	Melody        []LeadSheet      `yaml:"melody"`
	Chords        []ChordSheet     `yaml:"chords"`
	Key           string           `yaml:"key"`
	Tempo         float64          `yaml:"tempo"`
	TimeSignature g9.TimeSignature `yaml:"timesignature"`
	Title         string           `yaml:"title"`
}

func (s *Song) Load(songFile string) error {
	b, err := os.ReadFile(songFile)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(b, s)
	return err
}

func (s *Song) Initialize() error {
	return g9.NoteDuration.SetDuration(s.TimeSignature.Denominator, s.Tempo)
}

func (s *Song) ToStream() beep.Streamer {
	kse1 := karplusstrong.NewExtended(SampleRate, 0.1)
	lead := g9.NewGuitar(SampleRate, kse1)
	kse2 := karplusstrong.NewExtended(SampleRate, 0.05)
	rhythm := g9.NewGuitar(SampleRate, kse2)
	return s.Mixer(lead, rhythm)
}

func (s *Song) Mixer(lead, rhythm *g9.Guitar) beep.Streamer {
	leadStreamer := s.LeadGuitar(lead)
	rhythmStreamer := s.RhythmGuitar(rhythm)
	return beep.Mix(leadStreamer, rhythmStreamer)
}

func (s *Song) LeadGuitar(lead *g9.Guitar) beep.Streamer {
	var leadSequence []beep.Streamer
	for _, m := range s.Melody {
		dur := g9.NoteDuration[g9.NoteType(m.Duration)].Duration
		if m.Note == g9.Rest {
			leadSequence = append(leadSequence, lead.Silence(dur))
			continue
		}
		note := *g9.NoteMap[m.Note]
		freq := note[m.Octave]
		leadSequence = append(leadSequence, lead.Pluck(freq, dur))
	}
	return beep.Seq(leadSequence...)
}

func (s *Song) RhythmGuitar(rhythm *g9.Guitar) beep.Streamer {
	var chordSequence []beep.Streamer
	for _, c := range s.Chords {
		// First, get the duration
		cont := false
		dur := g9.NoteDuration[g9.NoteType(c.Duration)].Duration
		var notes []float64
		for _, n := range c.Chord {
			if n.Note == g9.Rest {
				chordSequence = append(chordSequence, rhythm.Silence(dur))
				cont = true
				break
			}
			note := *g9.NoteMap[n.Note]
			notes = append(notes, note[n.Octave])
		}
		if cont {
			// We must have had a Rest, so continue with the next chord
			continue
		}
		chordSequence = append(chordSequence, rhythm.Chord(notes, dur, 0.0195))
	}
	return beep.Seq(chordSequence...)
}

// Print writes a document to the terminal describing the song
func (s *Song) Print() {
	fmt.Printf("Title: %v\n", s.Title)
	fmt.Printf("Key: %v\n", s.Key)
	fmt.Printf("Tempo: %v\n", s.Tempo)
	fmt.Printf("Time Signature Numerator: %v\n", s.TimeSignature.Numerator)
	fmt.Printf("Time Signature Denominator: %v\n", s.TimeSignature.Denominator)
	fmt.Println("Melody:")
	fmt.Println()
	for _, m := range s.Melody {
		fmt.Printf("\tNote: %v, Octave: %v, Duration: %v\n", m.Note, m.Octave, m.Duration)
	}
	fmt.Println("Chords:")
	fmt.Println()
	for _, c := range s.Chords {
		for _, n := range c.Chord {
			fmt.Printf("\tChord Note: %v, Octave: %v\n", n.Note, n.Octave)
		}
		fmt.Printf("\tDuration: %v\n", c.Duration)
		fmt.Println()
	}
}
