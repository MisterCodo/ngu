package mansionsandmanagers

import (
	"embed"
	"image"
	"log"

	"github.com/MisterCodo/ngu/plugins/locations"
)

//go:embed data/*
var assets embed.FS

const (
	prettyName = "Mansions & Managers"
	uglyName   = "MansionsAndManagers"
)

type mansionsandmanagers struct {
	img  image.Image
	mask locations.Mask
}

func (p *mansionsandmanagers) Image() image.Image   { return p.img }
func (p *mansionsandmanagers) Mask() locations.Mask { return p.mask }
func (p *mansionsandmanagers) PrettyName() string   { return prettyName }
func (p *mansionsandmanagers) UglyName() string     { return uglyName }

func init() {
	img, err := locations.FileToImage(assets, "data/MansionsAndManagers.png")
	if err != nil {
		log.Fatalf("map image not found: %s", err.Error())
	}

	_, mask, err := locations.FileToMask(assets, "data/MansionsAndManagers.txt")
	if err != nil {
		log.Fatalf("map mask not found: %s", err.Error())
	}

	locations.Add(uglyName, &mansionsandmanagers{img: img, mask: mask})
}
