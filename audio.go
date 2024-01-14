package mypong

import (
	"bytes"
	"embed"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

const SAMPLERATE = 48000

//go:embed sounds/*
var sounds embed.FS // Intégrer un répertoire entier

var (
	// audio context
	audioContext *audio.Context = audio.NewContext(SAMPLERATE)
	// raw mp3 bytes for boing
	audioDataBoing []byte
	// player for rain
	audioPlayerRain *audio.Player
)

func init() {
	var err error
	// Load data for boieng
	audioDataBoing, err = sounds.ReadFile("sounds/boing.mp3")
	if err != nil {
		panic(err)
	}

	// create a single player for background
	data, err := sounds.ReadFile("sounds/chase.mp3")
	if err != nil {
		panic(err)
	}
	d, err := mp3.DecodeWithSampleRate(SAMPLERATE, bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
	audioPlayerRain, err = audioContext.NewPlayer(d)
	if err != nil {
		panic(err)
	}

	// Play sounds
	UpdateAudioBackground()
	// PlayAudioBoing()
}

// PlayAudioBoing creates a player and plays the boing file
func PlayAudioBoing() {

	d, err := mp3.DecodeWithSampleRate(SAMPLERATE, bytes.NewReader(audioDataBoing))
	if err != nil {
		panic(err)
	}

	p, err := audioContext.NewPlayer(d)
	if err != nil {
		panic(err)
	}
	if p != nil {
		p.Rewind()
		p.Play()
	}
}

// Plays rain sound in a loop, using a single player
// Should be called within the update loop.
func UpdateAudioBackground() {
	if audioPlayerRain != nil && !audioPlayerRain.IsPlaying() {
		audioPlayerRain.Rewind()
		audioPlayerRain.Play()
	}
}
