package candyland

import (
	"embed"
	"image"
	"log"

	"github.com/MisterCodo/ngu/plugins/locations"
)

//go:embed data/*
var assets embed.FS

const (
	prettyName = "Candy Land"
	uglyName   = "CandyLand"
)

type candyland struct {
	img  image.Image
	mask locations.Mask
}

func (p *candyland) Image() image.Image   { return p.img }
func (p *candyland) Mask() locations.Mask { return p.mask }
func (p *candyland) PrettyName() string   { return prettyName }
func (p *candyland) UglyName() string     { return uglyName }

func init() {
	img, err := locations.FileToImage(assets, "data/CandyLand.png")
	if err != nil {
		log.Fatalf("map image not found: %s", err.Error())
	}

	_, mask, err := locations.FileToMask(assets, "data/CandyLand.txt")
	if err != nil {
		log.Fatalf("map mask not found: %s", err.Error())
	}

	locations.Add(uglyName, &candyland{img: img, mask: mask})
}
