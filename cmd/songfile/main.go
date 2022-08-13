package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/faiface/beep"
	"github.com/faiface/beep/wav"
	"github.com/j3fflan3/arpeggiator/player"
)

var (
	songFile = flag.String("songFile", "./songs/Mary Had a Little Lamb.yaml", "the name of the yaml file that contains the song information")
)

func main() {
	flag.Parse()
	song := &player.Song{}
	err := song.Load(*songFile)
	if err != nil {
		log.Fatalf("could not load songFile %v", *songFile)
	}
	song.Print()
	if err = song.Initialize(); err == nil {
		s := song.ToStream()
		title := strings.ReplaceAll(song.Title, " ", "_")
		name := fmt.Sprintf("%v_%vbpm.wav", title, song.Tempo)
		f, err := os.Create(name)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Saving to %s...\n", name)
		if err = wav.Encode(f, s, beep.Format{
			SampleRate:  beep.SampleRate(player.SampleRate),
			NumChannels: 2,
			Precision:   3}); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Saved.")
		return
	}
	log.Fatalf("something sucked... %v", err)
}
