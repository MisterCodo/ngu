package fleshworld

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
	prettyName = "Flesh World"
	uglyName   = "FleshWorld"
)

type fleshworld struct{}

func (p *fleshworld) Image() image.Image   { return img }
func (p *fleshworld) Mask() locations.Mask { return mask }
func (p *fleshworld) PrettyName() string   { return prettyName }
func (p *fleshworld) UglyName() string     { return uglyName }

func init() {
	locations.Add(uglyName, func() locations.Location { return &fleshworld{} })

	var err error
	img, err = locations.FileToImage(assets, "data/FleshWorld.png")
	if err != nil {
		log.Fatalf("map image not found: %s", err.Error())
	}

	_, mask, err = locations.FileToMask(assets, "data/FleshWorld.txt")
	if err != nil {
		log.Fatalf("map mask not found: %s", err.Error())
	}
}
