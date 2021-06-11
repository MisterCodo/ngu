package productionbox

import (
	"embed"
	"image"
	"log"

	"github.com/MisterCodo/ngu/plugins/beacons"
)

//go:embed data/*
var assets embed.FS
var img image.Image

type productionbox struct{}

func (p *productionbox) Effect() []beacons.Effect {
	return []beacons.Effect{
		{X: -1, Y: 1, Gain: 30.0},
		{X: 0, Y: 1, Gain: 30.0},
		{X: 1, Y: 1, Gain: 30.0},
		{X: -1, Y: 0, Gain: 30.0},
		{X: 1, Y: 0, Gain: 30.0},
		{X: -1, Y: -1, Gain: 30.0},
		{X: 0, Y: -1, Gain: 30.0},
		{X: 1, Y: -1, Gain: 30.0},
	}
}

func (p *productionbox) Category() beacons.Category {
	return beacons.Production
}

func (p *productionbox) BType() beacons.BType {
	return beacons.Box
}

func (p *productionbox) Image() image.Image { return img }

func init() {
	beacons.Add("b", func() beacons.Beacon { return &productionbox{} })

	var err error
	img, err = beacons.FileToImage(assets, "data/ProductionBox.png")
	if err != nil {
		log.Fatalf("beacon image not found: %s", err.Error())
	}
}
