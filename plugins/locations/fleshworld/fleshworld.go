package fleshworld

import (
	"embed"
	"image"
	"log"

	"github.com/MisterCodo/ngu/plugins/locations"
)

//go:embed data/*
var assets embed.FS

const (
	prettyName = "Flesh World"
	uglyName   = "FleshWorld"
)

type fleshworld struct {
	img  image.Image
	mask locations.Mask
}

func (p *fleshworld) Image() image.Image   { return p.img }
func (p *fleshworld) Mask() locations.Mask { return p.mask }
func (p *fleshworld) PrettyName() string   { return prettyName }
func (p *fleshworld) UglyName() string     { return uglyName }

func init() {
	img, err := locations.FileToImage(assets, "data/FleshWorld.png")
	if err != nil {
		log.Fatalf("map image not found: %s", err.Error())
	}

	_, mask, err := locations.FileToMask(assets, "data/FleshWorld.txt")
	if err != nil {
		log.Fatalf("map mask not found: %s", err.Error())
	}

	locations.Add(uglyName, &fleshworld{img: img, mask: mask})
}
