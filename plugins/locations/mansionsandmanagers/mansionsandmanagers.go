package mansionsandmanagers

import (
	"embed"
	"image"
	"log"

	"github.com/MisterCodo/ngu/plugins/locations"
)

//go:embed data/*
var assets embed.FS
var img image.Image
var mask locations.Mask

const (
	prettyName = "Mansions & Managers"
	uglyName   = "MansionsAndManagers"
)

type mansionsandmanagers struct{}

func (p *mansionsandmanagers) Image() image.Image   { return img }
func (p *mansionsandmanagers) Mask() locations.Mask { return mask }
func (p *mansionsandmanagers) PrettyName() string   { return prettyName }
func (p *mansionsandmanagers) UglyName() string     { return uglyName }

func init() {
	locations.Add(uglyName, func() locations.Location { return &mansionsandmanagers{} })

	var err error
	img, err = locations.FileToImage(assets, "data/MansionsAndManagers.png")
	if err != nil {
		log.Fatalf("map image not found: %s", err.Error())
	}

	_, mask, err = locations.FileToMask(assets, "data/MansionsAndManagers.txt")
	if err != nil {
		log.Fatalf("map mask not found: %s", err.Error())
	}
}
