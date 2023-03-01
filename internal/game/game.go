package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	config "test-station/config/game"
)

type Game struct {
	input *Input
	cfg   *config.Config
	ship  *Ship
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.cfg.ScreenWidth / 2, g.cfg.ScreenHeight / 2
}

func NewGame() *Game {
	cfg := config.LoadConfig()

	return &Game{
		input: &Input{},
		cfg:   cfg,
		ship: NewShip(cfg.ScreenWidth, cfg.ScreenHeight),
	}
}

func (g *Game) Update() error {
	g.input.Update(g.ship)
	return nil
}

func (g *Game) Draw (screen *ebiten.Image) {
	g.ship.Draw(screen, g.cfg)
}