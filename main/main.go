package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/xavier268/mypong"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("MyPong by Xavier268")
	if err := ebiten.RunGame(mypong.NewPong()); err != nil {
		log.Fatal(err)
	}
}
