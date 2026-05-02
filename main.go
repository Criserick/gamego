package main

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
}

//logica de jogo
func (g *Game) Update() error {
	return nil
}

//objetos a serem desenhados na tela
func (g *Game) Draw(screen *ebiten.Image) {
	
}

func (g *Game) Layout( outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight	
}


func main() {

	g  := &Game{}

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err) 
	}
}