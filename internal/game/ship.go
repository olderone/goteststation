package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
	config "test-station/config/game"
)

type Ship struct {
	image  *ebiten.Image
	width  int
	height int
	x float64 // x坐标
	y float64 // y坐标
}

var img *ebiten.Image

func init() {
	dir,_ := os.Getwd()
	var err error
	img, _, err = ebitenutil.NewImageFromFile(filepath.Join(dir, "statics\\images\\640.png"))
	if err != nil {
		log.Fatal(err)
	}
}

func NewShip(screenWidth, screenHeight int) *Ship {

	width, height := img.Size()
	ship := &Ship{
		image:  img,
		width:  width,
		height: height,
		x: float64(screenWidth-width) / 2,
		y: float64(screenHeight - height),
	}
	return ship
}

func (ship *Ship) Draw(screen *ebiten.Image, cfg *config.Config) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(cfg.ScreenWidth-ship.width)/2, float64(cfg.ScreenHeight-ship.height))
	screen.DrawImage(ship.image, op)
}