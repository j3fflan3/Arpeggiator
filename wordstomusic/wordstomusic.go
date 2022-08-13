package wordstomusic

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/j3fflan3/arpeggiator/g9"
	"github.com/j3fflan3/arpeggiator/player"
	"gopkg.in/yaml.v3"
)

// Letter. A type representation of a member of the alphabet, e.g., "A"
type Letter string

// Alphabet. I thought explaining this might be insulting. So I'm not going to.
var Alphabet []Letter = []Letter{
	"a",
	"b",
	"c",
	"d",
	"e",
	"f",
	"g",
	"h",
	"i",
	"j",
	"k",
	"l",
	"m",
	"n",
	"o",
	"p",
	"q",
	"r",
	"s",
	"t",
	"u",
	"v",
	"w",
	"x",
	"y",
	"z",
}

type note struct {
	Note   string
	Octave int
}

var WordMap map[Letter]note = map[Letter]note{}

// Scale - Default Bm7 b5 +7, ascending. Note names from g9 package (constants)
var DemoScale []string

// DemoDuration - values must add up to the length of 1+ bars.
// Any durations that exceed the Arpeggio.Bars length will be ignored.
// String note duration values, e.g,. Quarter, Sixth, Eighth, etc.
var DemoDuration []g9.NoteType

// tempo in BPM (beats per minute).  Can be a fraction, e.g., 119.394
var tempo float64 = 160

// timeSignature.  The time signature of the song.
//
// The default values are 4/4
var timeSignature g9.TimeSignature = g9.TimeSignature{
	Numerator:   g9.Quarter,
	Denominator: g9.Quarter,
}

type (
	// Arpeggio contains the RhythmPattern float64 and number of Bars to repeat
	Arpeggio struct {
		// RhythmPattern - Sequence of note durations for the arpeggio
		RhythmPattern []float64
		// Bars - the number of bars to repeat.
		//
		// E.g., if bars is 4 and
		// the time signature is 4/4, it will be 4 bars of 4/4
		Bars int
	}
)

// Initialize the WordMap with nil g9.Note
func init() {
	for i := 0; i < len(Alphabet); i++ {
		WordMap[Alphabet[i]] = note{}
	}
	// sort the Alphabet for readability
	sort.Slice(Alphabet, func(i, j int) bool {
		return Alphabet[i] < Alphabet[j]
	})
	// Load default Duration
	g9.NoteDuration.SetDuration(timeSignature.Denominator, tempo)

	// Load the Demo Scale
	DemoScale = append(DemoScale, g9.BNatural)
	//DemoScale = append(DemoScale, g9.CSharp)
	DemoScale = append(DemoScale, g9.DNatural)
	DemoScale = append(DemoScale, g9.ENatural)
	//DemoScale = append(DemoScale, g9.FNatural)
	DemoScale = append(DemoScale, g9.FSharp)
	DemoScale = append(DemoScale, g9.GNatural)
	//DemoScale = append(DemoScale, g9.ANatural)
	DemoScale = append(DemoScale, g9.ASharp)

	DemoDuration = append(DemoDuration, g9.Quarter)
	DemoDuration = append(DemoDuration, g9.Quarter)
	DemoDuration = append(DemoDuration, g9.Sixth)
	DemoDuration = append(DemoDuration, g9.Sixth)
	DemoDuration = append(DemoDuration, g9.Sixth)

}

func SetTimeSignature(numerator, denominator g9.NoteType) {
	timeSignature.Numerator = numerator
	timeSignature.Denominator = denominator
}
func SetDuration(bpm float64) {
	tempo = bpm
	g9.NoteDuration.SetDuration(timeSignature.Denominator, tempo)
}

// DemoArpeggio
//
// Builds and writes a demo wav file from the phrase "hello world" and
// a Bm7 scale.
//
// Note that len(demoPhrase) equals the total note duration of 2 bars
// with our rhythm pattern.
func DemoArpeggio() {
	SetDuration(161)
	demoPhrase := []string{
		"h", "e", "l", "l", "o", "w", "o", "r", "l", "d",
	}
	octaves := []int{3, 2, 3}
	filepath := "../arpeggios/"
	span := 2
	bars := 16
	err := BuildArpeggio(demoPhrase, DemoScale, DemoDuration, octaves, span, bars, "Bm", filepath, "helloworld")
	if err != nil {
		fmt.Println(err)
	}
}

// BuildArpeggio
//
// phrase []string.
// The phrase to turn into musical notes
//
// sequence []g9.NoteType.
// The rhythm pattern, an array of note lengths (duration) per bar
//
// scale []string.
// The slice of string to use as the scale for our sequence
//
// octave []int.
// The slice of octaves to apply to Alphabet letters based on scale length
//
// span int.
// The number of bars the sequence spans.  For instance, the "helloworld" demo sequence pattern
// is quarter-quarter-sixth-sixth-sixth and the phrase is 10 characters long, so the span would be
// 2 bars in 4/4 time.
//
// bars int.
// The number of total bars of the arpeggio to play/repeat.
//
// key string.
// Key signature of the song (for descriptive purposes only)
//
// filename string.
// The name of the yaml file to write the arpeggio to.  Should include the full or relative file path.
func BuildArpeggio(phrase, scale []string, sequence []g9.NoteType, octaves []int, span, bars int, key, filepath, title string) error {
	// min/max length of our scale
	i, o := 0, 0
	iMax, oMax := len(scale), len(octaves)
	// Assign alphabet letters a note value
	for _, k := range Alphabet {
		octave := octaves[o]
		fmt.Println(i, k, octave)
		WordMap[k] = note{Note: scale[i], Octave: octave}
		if i == iMax-1 {
			i = 0
			if o == oMax-1 {
				o = 0
			} else {
				o++
			}
		} else {
			i++
		}
	}
	PrintWordMap()
	leadSheet := []player.LeadSheet{}
	// index of the sequence
	idx := 0
	// max index of sequence before starting over
	idxMax := len(sequence)
	for n := span - 1; n < bars; n = n + span {
		for _, l := range phrase {
			nt := WordMap[Letter(l)]
			ls := player.LeadSheet{
				Note:     nt.Note,
				Octave:   nt.Octave,
				Duration: string(sequence[idx]),
			}
			leadSheet = append(leadSheet, ls)
			if idx == idxMax-1 {
				idx = 0
			} else {
				idx++
			}
		}
	}
	song := player.Song{
		Title:         title,
		TimeSignature: timeSignature,
		Tempo:         tempo,
		Key:           key,
		Melody:        leadSheet,
	}
	return SaveArpeggio(song, filepath, title)
}

// SaveArpeggio
//
// @song player.Song.
// Hydrated Song struct to save to yaml.
//
// @path string.
// The file path to save the song yaml to.
//
// @title string.
// The Title of the song.
//
// Saves the song into a yaml file with the given path and title.  The yaml file can
// be passed as an argument to the player to generate a .WAV file.
func SaveArpeggio(song player.Song, path, title string) error {
	ext := ".yaml"
	filename := title + ext
	f, err := os.Create(filepath.Join(path, filename))
	if err != nil {
		return err
	}
	b, err := yaml.Marshal(song)
	if err != nil {
		return err
	}
	_, err = f.Write(b)
	if err != nil {
		return err
	}
	return f.Close()
}

// PrintWordMap
//
// Prints the WordMap, you dunce cap.
func PrintWordMap() {
	for _, k := range Alphabet {
		v := WordMap[k]
		fmt.Printf("Letter: %v - Note: %v\n", k, v)
	}
}
