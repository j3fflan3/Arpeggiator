package wordstomusic

import (
	"fmt"

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

var WordMap map[Letter]g9.Note = map[Letter]g9.Note{}

// Initialize the WordMap with nil g9.Note
func init() {
	for i := 0; i < len(Alphabet); i++ {
		WordMap[Alphabet[i]] = nil
	}
}

func PrintWordMap() {
	for k, v := range WordMap {
		fmt.Printf("Letter: %v - Note: %v\n", k, v)
	}
}
