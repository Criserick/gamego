package main

import (
	"mygame/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	g  := &game.Game{}

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err) 
	}
}