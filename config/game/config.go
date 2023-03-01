package config


import (
	"encoding/json"
	"image/color"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	BulletWidth  int 		`json:"bulletWidth"`
	BulletHeight int 		`json:"bulletHeight"`
	BulletSpeedFactor int 	`json:"bulletSpeedFactor"`
	MaxBulletNum      int   `json:"maxBulletNum"`
	BulletInterval    int64 `json:"bulletInterval"`
	BulletColor  color.RGBA `json:"bulletColor"`
	ScreenWidth  int        `json:"screenWidth"`
	ScreenHeight int        `json:"screenHeight"`
	Title        string     `json:"title"`
	BgColor      color.RGBA `json:"bgColor"`
	ShipSpeedFactor int    `json:"shipSpeedFactor"`
	AlienSpeedFactor int `json:"alienSpeedFactor"`
	TitleFontSize int `json:"titleFontSize"`
	FontSize int `json:"fontSize"`
	SmallFontSize int `json:"smallFontSize"`
}

func LoadConfig() *Config {
	dir,_ := os.Getwd()
	f, err := os.Open(filepath.Join(dir, "config/game/config.json"))
	if err != nil {
		log.Fatalf("os.Open failed: %v\n", err)
	}

	var cfg Config
	err = json.NewDecoder(f).Decode(&cfg)
	if err != nil {
		log.Fatalf("json.Decode failed: %v\n", err)
	}

	return &cfg
}

