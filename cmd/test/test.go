package main

import (
	"log"
	"test-station/internal/game"

	"github.com/hajimehoshi/ebiten/v2"
)


func main() {

	gm := game.NewGame()

	if err := ebiten.RunGame(gm); err != nil {
		log.Fatal(err)
	}
}