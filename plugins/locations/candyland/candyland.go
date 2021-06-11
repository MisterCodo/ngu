package candyland

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
	prettyName = "Candy Land"
	uglyName   = "CandyLand"
)

type candyland struct{}

func (p *candyland) Image() image.Image   { return img }
func (p *candyland) Mask() locations.Mask { return mask }
func (p *candyland) PrettyName() string   { return prettyName }
func (p *candyland) UglyName() string     { return uglyName }

func init() {
	locations.Add(uglyName, func() locations.Location { return &candyland{} })

	var err error
	img, err = locations.FileToImage(assets, "data/CandyLand.png")
	if err != nil {
		log.Fatalf("map image not found: %s", err.Error())
	}

	_, mask, err = locations.FileToMask(assets, "data/CandyLand.txt")
	if err != nil {
		log.Fatalf("map mask not found: %s", err.Error())
	}
}
