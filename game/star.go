package game

import (
	"math/rand"
	"mygame/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Star struct {
	image    *ebiten.Image
	speed    float64
	position Vector
}

func NewStar() *Star {
	image := assets.StarsSprites[rand.Intn(len(assets.StarsSprites))]

	speed := (rand.Float64() * 3)

	position := Vector{
		X: rand.Float64() * ScreenWidth,
		Y: -100,
	}

	return &Star{
		image:    image,
		speed:    speed,
		position: position,
	}
}

func (s *Star) Update() {
	s.position.Y += s.speed
}

func (s *Star) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	//posicionamento da imagem
	op.GeoM.Translate(s.position.X, s.position.Y)

	//desenho da imagem
	screen.DrawImage(s.image, op)
}