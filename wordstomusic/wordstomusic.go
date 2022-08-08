package wordstomusic

import (
	"fmt"
	"sort"

	"github.com/j3fflan3/arpeggiator/g9"
)

type Letter string

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

var WordMap map[Letter]*g9.Note = map[Letter]*g9.Note{}

// Scale - Default Bm7 b5 +7, ascending
var DemoScale []*g9.Note

// DemoDuration - values must add up to the length of 1+ bars.
// Any durations that exceed the Arpeggio.Bars length will be ignored.
var DemoDuration []float64

// Tempo in beats per minute.  Can be a fraction, e.g., 119.394
var Tempo float64 = 120

// TimeSignature.  The time signature of the song, e.g., 4/4
var TimeSignature g9.TimeSignature = g9.TimeSignature{
	Numerator:   g9.Quarter,
	Denominator: g9.Quarter,
}

type (
	// Arpeggio contains the RhythmPattern float64 and number of Bars to repeat
	Arpeggio struct {
		// RhythmPattern - Sequence of note durations for the arpeggio
		RhythmPattern []float64
		// Bars - the number of bars to repeat.  E.g., if bars is 4 and
		// the time signature is 4/4, it will be 4 bars of 4/4
		Bars int
	}
)

// Initialize the WordMap with nil g9.Note
func init() {
	for i := 0; i < len(Alphabet); i++ {
		WordMap[Alphabet[i]] = nil
	}
	// sort the Alphabet for readability
	sort.Slice(Alphabet, func(i, j int) bool {
		return Alphabet[i] < Alphabet[j]
	})
	// Load default Duration
	g9.NoteDuration.SetDuration(TimeSignature.Denominator, Tempo)

	// Load the Demo Scale
	DemoScale = append(DemoScale, g9.NoteMap[g9.BNatural])
	DemoScale = append(DemoScale, g9.NoteMap[g9.CSharp])
	DemoScale = append(DemoScale, g9.NoteMap[g9.DNatural])
	DemoScale = append(DemoScale, g9.NoteMap[g9.ENatural])
	// DemoScale = append(DemoScale, g9.NoteMap[g9.FNatural])
	DemoScale = append(DemoScale, g9.NoteMap[g9.FSharp])
	DemoScale = append(DemoScale, g9.NoteMap[g9.GNatural])
	// DemoScale = append(DemoScale, g9.NoteMap[g9.ANatural])
	DemoScale = append(DemoScale, g9.NoteMap[g9.ASharp])

	DemoDuration = append(DemoDuration, g9.NoteDuration[g9.Quarter].Duration)
	DemoDuration = append(DemoDuration, g9.NoteDuration[g9.Quarter].Duration)
	DemoDuration = append(DemoDuration, g9.NoteDuration[g9.Sixth].Duration)
	DemoDuration = append(DemoDuration, g9.NoteDuration[g9.Sixth].Duration)
	DemoDuration = append(DemoDuration, g9.NoteDuration[g9.Sixth].Duration)

}

// DemoArpeggio
//
// Builds and writes a demo wav file from the phrase "hello world" and
// a Bm7 scale.
//
// Note that len(demoPhrase) equals the total note duration of 2 bars
// with our rhythm pattern.
func DemoArpeggio() {
	demoPhrase := []string{
		"h", "e", "l", "l", "o", "w", "o", "r", "l", "d",
	}
	octaves := []int{2, 1, 3}
	filename := "helloworld"
	BuildArpeggio(demoPhrase, DemoDuration, DemoScale, octaves, filename)
}

// BuildArpeggio
//
// 		 phrase []string - The phrase to turn into musical notes
//
//		 sequence []float64 - The rhythm pattern, an array of note lengths (duration) per bar
//
//		 scale []*g9.Note - The slice of scale to use in our sequence
//
//		 octave []int - The slice of ocatves to apply to Alphabet letters based on scale length
func BuildArpeggio(phrase []string, sequence []float64, scale []*g9.Note, octaves []int, filename string) error {
	// min/max length of our scale
	i, o := 0, 0
	iMax, oMax := len(scale), len(octaves)
	// Assign alphabet letters a note value
	for _, k := range Alphabet {
		fmt.Println(i, k, o)
		WordMap[k] = scale[i]
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
	return nil
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
