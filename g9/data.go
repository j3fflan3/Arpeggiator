package g9

import "fmt"

// Note types map[8va]Hz
type Note map[int]float64
type NoteType string

const (
	// Note names
	Cnatural string = "C"
	Csharp   string = "C#"
	Dflat    string = "Db"
	Dnatural string = "D"
	Dsharp   string = "D#"
	Eflat    string = "Eb"
	Enatural string = "E"
	Fnatural string = "F"
	Fsharp   string = "F#"
	Gflat    string = "Gb"
	Gnatural string = "G"
	Gsharp   string = "G#"
	Aflat    string = "Ab"
	Anatural string = "A"
	Asharp   string = "A#"
	Bflat    string = "Bb"
	Bnatural string = "B"
	Rest     string = "rest"
	// Note Duration Types
	WholeDotted        NoteType = "wholedotted"        // whole note * 1.5
	Whole              NoteType = "whole"              // Whole note
	HalfDotted         NoteType = "halfdotted"         // half note * 1.5
	Half               NoteType = "half"               // Half note
	QuarterDotted      NoteType = "quarterdotted"      // quarter note * 1.5
	Quarter            NoteType = "quarter"            // Quarter note
	Fifth              NoteType = "fifth"              // Quintuplets... five beats per measure where time signature is 4/4
	Sixth              NoteType = "sixth"              // represents quarter note triplet
	EighthDotted       NoteType = "eighthdotted"       // eighth note * 1.5
	Eighth             NoteType = "eighth"             // eighth note
	Tenth              NoteType = "tenth"              // Quintuplets for eighth notes, e.g. where time signature is 4/4
	Twelfth            NoteType = "twelfth"            // represents eighth note triplet
	SixteenthDotted    NoteType = "sixteenthdotted"    // sixteenth note * 1.5
	Sixteenth          NoteType = "sixteenth"          // sixteenth note
	Twentieth          NoteType = "twentieth"          // quintuplets for sixteenth notes, e.g., when time signature is 4/4
	Twentyfourth       NoteType = "twentyfourth"       // represents sixteenth note triplet
	ThirtysecondDotted NoteType = "thirtyseconddotted" // thirtysecond note * 1.5
	Thirtysecond       NoteType = "thirtysecond"       // thirtysecond note
	// Max Denominator
	MxD             int     = 64
	SecondsInMinute float64 = 60
)

var NoteMap map[string]*Note = map[string]*Note{
	Cnatural: &C,
	Csharp:   &Db,
	Dflat:    &Db,
	Dnatural: &D,
	Dsharp:   &Eb,
	Eflat:    &Eb,
	Enatural: &E,
	Fnatural: &F,
	Fsharp:   &Gb,
	Gflat:    &Gb,
	Gnatural: &G,
	Gsharp:   &Ab,
	Aflat:    &Ab,
	Anatural: &A,
	Asharp:   &Bb,
	Bflat:    &Bb,
	Bnatural: &B,
}

// Note octaves and frequencies
var C = Note{
	2: 65.41,
	3: 130.8,
	4: 261.6,
	5: 523.3,
	6: 1047,
}

var Db = Note{
	1: 34.65,
	2: 69.3,
	3: 138.6,
	4: 277.2,
	5: 554.4,
	6: 1109,
}

var D = Note{
	1: 36.71,
	2: 73.42,
	3: 146.8,
	4: 293.7,
	5: 587.3,
	6: 1175,
}

var Eb = Note{
	1: 38.89,
	2: 77.78,
	3: 155.6,
	4: 311.1,
	5: 622.3,
	6: 1245,
}

var E = Note{
	1: 41.2,
	2: 82.41,
	3: 164.8,
	4: 329.6,
	5: 659.3,
	6: 1319,
}

var F = Note{
	1: 43.65,
	2: 87.31,
	3: 174.6,
	4: 349.2,
	5: 698.5,
}

var Gb = Note{
	1: 46.25,
	2: 92.5,
	3: 185,
	4: 370,
	5: 740,
}
var G = Note{
	1: 49,
	2: 98,
	3: 196,
	4: 392,
	5: 784,
}
var Ab = Note{
	1: 51.91,
	2: 103.8,
	3: 207.7,
	4: 415.3,
	5: 830.6,
}
var A = Note{
	1: 55,
	2: 110,
	3: 220,
	4: 440,
	5: 880,
}

var Bb = Note{
	1: 58.27,
	2: 116.5,
	3: 233.1,
	4: 466.2,
	5: 932.3,
}
var B = Note{
	1: 61.74,
	2: 123.5,
	3: 246.9,
	4: 493.9,
	5: 987.8,
}

type Duration struct {
	Duration      float64 // Duration of the note type in seconds
	Denominator   bool    // false for non-denominator values, true for denominator values.
	NumericValue  float64 // The fractional value of the numeric representation of the note type.
	Divisor       float64 // Divisor where the dividend is 1
	IsDenominator bool    // True indicates this note type is the denominator of the time signature.
}

type TNoteDuration map[NoteType]Duration

// NoteDuration contains the initialized values with 0 duration for each note.
// Actual durations are calculated in SetNoteDurations given a time signature
// denominator and a tempo (BPM).
//  time signature:   numerator   e.g.:   3
//                    ---------           -
//                    denominator         4
// Maximum denominator for the current iteration
// of this application is 32. Denominators must be exponents 2¹ - 2⁵
var NoteDuration TNoteDuration = TNoteDuration{
	WholeDotted: {
		Duration:      0,
		Denominator:   false, // True means this note can be used as a denominator
		NumericValue:  1.5,
		Divisor:       0.666666666666667,
		IsDenominator: false, // True means this note is the current denominator
	}, // whole note * 1.5
	Whole: {
		Duration:      0,
		Denominator:   true,
		NumericValue:  1,
		Divisor:       1,
		IsDenominator: false,
	}, // whole
	HalfDotted: {
		Duration:      0,
		Denominator:   false, // True means this note can be used as a denominator
		NumericValue:  0.75,
		Divisor:       1.333333333333333,
		IsDenominator: false,
	}, // half note * 1.5
	Half: {
		Duration:      0,
		Denominator:   true,
		NumericValue:  0.5,
		Divisor:       2,
		IsDenominator: false,
	}, // half
	QuarterDotted: {
		Duration:      0,
		Denominator:   false, // True means this note can be used as a denominator
		NumericValue:  0.375,
		Divisor:       2.666666666666667,
		IsDenominator: false,
	}, // quarter note * 1.5
	Quarter: {
		Duration:      0, // 60/bpm = beat duration. So, bpm=140, duration of denominator would be .42857
		Denominator:   true,
		NumericValue:  0.25,
		Divisor:       4,
		IsDenominator: false,
	}, // quarter
	Fifth: {
		Duration:      0,
		Denominator:   false,
		NumericValue:  0.2,
		Divisor:       5,
		IsDenominator: false,
	},
	Sixth: {
		Duration:      0,
		Denominator:   false, // True means this note can be used as a denominator
		NumericValue:  0.166666666666667,
		Divisor:       6,
		IsDenominator: false,
	}, // represents quarter note triplet
	EighthDotted: {
		Duration:      0,
		Denominator:   false, // True means this note can be used as a denominator
		NumericValue:  0.1875,
		Divisor:       5.333333333333333,
		IsDenominator: false,
	}, // eighth note * 1.5
	Eighth: {
		Duration:      0,
		Denominator:   true,
		NumericValue:  0.125,
		Divisor:       8,
		IsDenominator: false,
	}, // eighth
	Tenth: {
		Duration:      0,
		Denominator:   false,
		NumericValue:  0.1,
		Divisor:       10,
		IsDenominator: false,
	},
	Twelfth: {
		Duration:      0,
		Denominator:   false, // True means this note can be used as a denominator
		NumericValue:  0.083333333333333,
		Divisor:       12,
		IsDenominator: false,
	}, // represents eighth note triplet
	SixteenthDotted: {
		Duration:      0,
		Denominator:   false, // True means this note can be used as a denominator
		NumericValue:  0.09375,
		Divisor:       10.666666666666667,
		IsDenominator: false,
	}, // sixteenth note * 1.5
	Sixteenth: {
		Duration:      0,
		Denominator:   true,
		NumericValue:  0.0625,
		Divisor:       16,
		IsDenominator: false,
	}, // sixteenth
	Twentieth: {
		Duration:      0,
		Denominator:   false,
		NumericValue:  .05,
		Divisor:       20,
		IsDenominator: false,
	},
	Twentyfourth: {
		Duration:      0,
		Denominator:   false, // True means this note can be used as a denominator
		NumericValue:  0.041666666666667,
		Divisor:       24,
		IsDenominator: false,
	}, // represents sixteenth note triplet
	ThirtysecondDotted: {
		Duration:      0,
		Denominator:   false, // True means this note can be used as a denominator
		NumericValue:  0.046875,
		Divisor:       21.333333333333333,
		IsDenominator: false,
	}, // thirtysecond note * 1.5
	Thirtysecond: {
		Duration:      0,
		Denominator:   true,
		NumericValue:  0.03125,
		Divisor:       32,
		IsDenominator: false,
	}, // thirtysecond
}

func (n Note) Print(octave int) {
	freq, ok := n[octave]
	if !ok {
		fmt.Printf("octave %d note found\n", octave)
		return
	}
	fmt.Println(freq)
}

/*
SetDuration sets the duration in the NoteDuration map based on the
time signature denominator and the tempo (BPM).  It returns an error if
the denominator is an unknown note or if the denominator's numeric value
is not even.

Tempo and Note Duration Calculation

To get the duration of the denominator (lower number) in a time signature, you need to perform the following calculation:

denominatorDuration = 60 / bpm

To extrapolate the duration of other note types, you need to use this calculation. If note type is not a denominator (eg dotted or triplet), calculate the numeric divisor by dividing the notes decimal representation (e.g., dotted whole = 1.5) by 1.

Example: To calculate a dotted whole note divisor, divide the dotted whole's numeric representation by 1:
1.5 / 1 = 0.666666666666667

Now we need a calculation to determine the target note's duration.

      WHOTE NOTE                     Denominator
duration = ((SecondsInMinute / bpm) * divisor) / targetNoteDivisor

So, given a bpm of 140 and a time signature of 4/4, what is the duration in seconds of a dotted Whole note?
((60/140) * 4) / 0.666666666666667 = 2.5714285714285703 seconds
^^^^^^^^^^^^^^   ^^^^^^^^^^^^^^^^^   ^^^^^^^^^^^^^^^^^^^^^^^^^^
Whole Note      Dotted Whole Divisor Dotted Whole Note Duration
*/
func (nd TNoteDuration) SetDuration(denominator NoteType, bpm float64) error {
	if _, ok := nd[denominator]; !ok {
		return fmt.Errorf("unexpected note type: %v", denominator)
	}
	d := nd[denominator]
	if !d.Denominator {
		return fmt.Errorf("the maximum dividend %v must be divisible without remainder by denominator %v", MxD, denominator)
	}
	d.IsDenominator = true
	denominatorDuration := 60 / bpm
	denominatorDivisor := d.Divisor
	d.Duration = denominatorDuration

	nd[denominator] = d
	for k, v := range nd {
		if k != denominator {
			v.Duration = calcNoteDuration(bpm, denominatorDivisor, v.Divisor)
			if _, ok := nd[k]; !ok {
				return fmt.Errorf("note type %v not found", k)
			}
			nd[k] = v
		}
	}
	return nil
}

/*
calcNoteDuration
@bpm float64 - Beats Per Minute
@divisor float64 - The denominator of the time signature, e.g., in 3/4 time the numerator is 3 and the denominator is 4

*/
func calcNoteDuration(bpm, divisor, targetNoteDivisor float64) float64 {
	// denominator duration = 60 seconds / bpm
	return ((SecondsInMinute / bpm) * divisor) / targetNoteDivisor
}
