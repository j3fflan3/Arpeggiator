# Arpeggiator

Arpeggiator is a configurable arpeggiator written in Go

g9 is a package for modeling 9-string synth guitar with a note range of C#1 to E6.

The arpeggiator features a "words" to music where you can assign custom scales to the alphabet (spanning multiple octaves). When a user writes a message in english, it will take the letters of each word and assign a note. It is configurable to assign which octave should be covered on each section of the alphabet the custom scale covers.

Note sequences and durations can be specified manually, or again via words using Morse Code. You can configure which note durations to use with Morse Code, for example Quarter Notes for dots . and Half Notes for dashes -

Credits:

### This is built on the work done by

## Michal Štrba https://github.com/faiface/beep

## Timur Iskhakov https://timiskhakov.github.io/posts/programming-guitar-music
