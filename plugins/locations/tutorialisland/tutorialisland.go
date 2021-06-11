package tutorialisland

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
	prettyName = "Tutorial Island"
	uglyName   = "TutorialIsland"
)

type tutorialisland struct{}

func (p *tutorialisland) Image() image.Image   { return img }
func (p *tutorialisland) Mask() locations.Mask { return mask }
func (p *tutorialisland) PrettyName() string   { return prettyName }
func (p *tutorialisland) UglyName() string     { return uglyName }

func init() {
	locations.Add(uglyName, func() locations.Location { return &tutorialisland{} })

	var err error
	img, err = locations.FileToImage(assets, "data/TutorialIsland.png")
	if err != nil {
		log.Fatalf("map image not found: %s", err.Error())
	}

	_, mask, err = locations.FileToMask(assets, "data/TutorialIsland.txt")
	if err != nil {
		log.Fatalf("map mask not found: %s", err.Error())
	}
}
