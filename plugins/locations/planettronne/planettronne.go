package planettronne

import (
	"embed"
	"image"
	"log"

	"github.com/MisterCodo/ngu/plugins/locations"
)

//go:embed data/*
var assets embed.FS

const (
	prettyName = "Planet Tronne"
	uglyName   = "PlanetTronne"
)

type planettronne struct {
	img  image.Image
	mask locations.Mask
}

func (p *planettronne) Image() image.Image   { return p.img }
func (p *planettronne) Mask() locations.Mask { return p.mask }
func (p *planettronne) PrettyName() string   { return prettyName }
func (p *planettronne) UglyName() string     { return uglyName }

func init() {
	img, err := locations.FileToImage(assets, "data/PlanetTronne.png")
	if err != nil {
		log.Fatalf("map image not found: %s", err.Error())
	}

	_, mask, err := locations.FileToMask(assets, "data/PlanetTronne.txt")
	if err != nil {
		log.Fatalf("map mask not found: %s", err.Error())
	}

	locations.Add(uglyName, &planettronne{img: img, mask: mask})
}
