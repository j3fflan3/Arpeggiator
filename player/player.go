package player

import (
	"fmt"

	"github.com/faiface/beep"
	"github.com/j3fflan3/arpeggiator/g9"
	"gopkg.in/yaml.v3"
)

var Streams []beep.Streamer
var Sequence []NoteEvent

var CurrentSong Song

type NoteEvent struct {
	Note     float64 `yaml:"note"`
	Duration float64 `yaml:"duration"`
}

type ChordEvent struct {
	Notes []NoteEvent
	Name  string
}

type Song struct {
	Chords []ChordEvent `yaml:"chords"`
	Notes  []NoteEvent  `yaml:"notes"`
	Key    string       `yaml:"key"`
	Tempo  float64      `yaml:"tempo"`
	g9.TimeSignature
	Title string `yaml:"title"`
}

func LoadSong(songFile string) {

}

func PrintTestSongYaml() {
	song := Song{}
	chords := []ChordEvent{}
	song.Key = "C"
	song.Tempo = 110
	song.TimeSignature = g9.TimeSignature{Numerator: g9.Quarter, Denominator: g9.Quarter}
	song.Title = "Mary Had a Little Lamb"

	g9.NoteDuration.SetDuration(g9.Quarter, 110)
	song.Chords = chords
	song.Notes = maryNotes()

	b, err := yaml.Marshal(song)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}

func maryNotes() []NoteEvent {
	e4 := g9.E[4]
	d4 := g9.D[4]
	c4 := g9.C[4]
	g4 := g9.G[4]
	fmt.Println(e4, d4, c4, g4)
	d8th := g9.NoteDuration[g9.Eighth].Duration
	d4th := g9.NoteDuration[g9.Quarter].Duration
	whole := g9.NoteDuration[g9.Whole].Duration
	return []NoteEvent{
		{Note: g9.E[4], Duration: d8th},
		{Note: g9.D[4], Duration: d8th},
		{Note: g9.C[4], Duration: d8th},
		{Note: g9.D[4], Duration: d8th},
		{Note: g9.E[4], Duration: d8th},
		{Note: g9.E[4], Duration: d8th},
		{Note: g9.E[4], Duration: d4th},
		{Note: g9.E[2], Duration: d8th},
		{Note: g9.E[2], Duration: d8th},
		{Note: g9.E[2], Duration: d4th},
		{Note: g9.E[3], Duration: d8th},
		{Note: g9.E[5], Duration: d8th},
		{Note: g9.E[5], Duration: d4th},
		{Note: g9.E[4], Duration: d8th},
		{Note: g9.D[4], Duration: d8th},
		{Note: g9.C[4], Duration: d8th},
		{Note: g9.D[4], Duration: d8th},
		{Note: g9.E[4], Duration: d8th},
		{Note: g9.E[4], Duration: d8th},
		{Note: g9.E[4], Duration: d8th},
		{Note: g9.E[4], Duration: d8th},
		{Note: g9.D[4], Duration: d8th},
		{Note: g9.D[4], Duration: d8th},
		{Note: g9.E[4], Duration: d8th},
		{Note: g9.D[4], Duration: d8th},
		{Note: g9.C[4], Duration: whole},
	}
}
