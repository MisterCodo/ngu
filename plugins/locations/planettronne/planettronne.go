package planettronne

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
	prettyName = "Planet Tronne"
	uglyName   = "PlanetTronne"
)

type planettronne struct{}

func (p *planettronne) Image() image.Image   { return img }
func (p *planettronne) Mask() locations.Mask { return mask }
func (p *planettronne) PrettyName() string   { return prettyName }
func (p *planettronne) UglyName() string     { return uglyName }

func init() {
	locations.Add(uglyName, func() locations.Location { return &planettronne{} })

	var err error
	img, err = locations.FileToImage(assets, "data/PlanetTronne.png")
	if err != nil {
		log.Fatalf("map image not found: %s", err.Error())
	}

	_, mask, err = locations.FileToMask(assets, "data/PlanetTronne.txt")
	if err != nil {
		log.Fatalf("map mask not found: %s", err.Error())
	}
}
