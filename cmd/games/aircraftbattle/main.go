package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"test-station/internal/cbg"
)

func main() {
	g := cbg.NewGame()
	ebiten.SetWindowSize(g.Cfg.ScreenWidth, g.Cfg.ScreenHeight)
	ebiten.SetWindowTitle(g.Cfg.Title)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

