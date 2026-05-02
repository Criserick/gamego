package game

import (
	"mygame/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	image *ebiten.Image
	position Vector
	game *Game
	laserLoadingTimer *Timer
}

func NewPlayer(game *Game) *Player {
	image := assets.PlayerSprite

	bounds :=image.Bounds()
	halfw := float64(bounds.Dx()) / 2

	position := Vector{
		X: (ScreenWidth / 2) - halfw,
		Y: 500,
	}

	return &Player{
		image: image,
		game : game,
		position: position,
		laserLoadingTimer: NewTimer(12),
	}
}

func (p *Player) Update() { 

	speed := 5.0

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.position.X -= speed
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) { 
		p.position.X += speed
	}

	p.laserLoadingTimer.Update()

	if ebiten.IsKeyPressed(ebiten.KeySpace) && p.laserLoadingTimer.IsReady() {
	p.laserLoadingTimer.Reset()
	
	bounds := p.image.Bounds()

	halfw := float64(bounds.Dx()) / 2 
	halfh := float64(bounds.Dy()) / 2

	spawnPos := Vector{
	X: p.position.X + halfw,
	Y: p.position.Y + halfh/2,
	}

	laser := NewLaser(spawnPos)
	p.game.AddLaser(laser)
		
	}
 }

func (p *Player) Draw(screen *ebiten.Image){	
	op := &ebiten.DrawImageOptions{}

	//posicionamento da imagem
	op.GeoM.Translate(p.position.X, p.position.Y)

	//desenho da imagem
	screen.DrawImage(p.image, op)
	}

