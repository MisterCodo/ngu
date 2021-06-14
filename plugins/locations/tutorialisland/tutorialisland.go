package tutorialisland

import (
	"embed"
	"image"
	"log"

	"github.com/MisterCodo/ngu/plugins/locations"
)

//go:embed data/*
var assets embed.FS

const (
	prettyName = "Tutorial Island"
	uglyName   = "TutorialIsland"
)

type tutorialisland struct {
	img  image.Image
	mask locations.Mask
}

func (p *tutorialisland) Image() image.Image   { return p.img }
func (p *tutorialisland) Mask() locations.Mask { return p.mask }
func (p *tutorialisland) PrettyName() string   { return prettyName }
func (p *tutorialisland) UglyName() string     { return uglyName }

func init() {
	img, err := locations.FileToImage(assets, "data/TutorialIsland.png")
	if err != nil {
		log.Fatalf("map image not found: %s", err.Error())
	}

	_, mask, err := locations.FileToMask(assets, "data/TutorialIsland.txt")
	if err != nil {
		log.Fatalf("map mask not found: %s", err.Error())
	}

	locations.Add(uglyName, &tutorialisland{img: img, mask: mask})
}
