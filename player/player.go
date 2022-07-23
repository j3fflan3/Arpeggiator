package player

import "github.com/faiface/beep"

var Streams []beep.Streamer
var Sequence []NoteEvent

var CurrentSong Song

type NoteEvent struct {
	Note     string  `yaml:"note"`
	Octave   int     `yaml:"octave"`
	Duration float64 `yaml:"duration"`
}

type ChordEvent struct {
	Notes []NoteEvent
	Name  string
}

type Song struct {
	Chords []string    `yaml:"chords"`
	Notes  []NoteEvent `yaml:"notes"`
	Key    string      `yaml:"key"`
	Tempo  float64     `yaml:"tempo"`
	Title  string      `yaml:"title"`
}

func LoadSong(songFile string) {

}
