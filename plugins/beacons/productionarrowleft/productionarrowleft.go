package productionarrowleft

import (
	"embed"
	"image"
	"log"

	"github.com/MisterCodo/ngu/plugins/beacons"
)

//go:embed data/*
var assets embed.FS
var img image.Image

type productionarrowleft struct{}

func (p *productionarrowleft) Effect() []beacons.Effect {
	return []beacons.Effect{
		{X: -3, Y: 2, Gain: 22.0},

		{X: -4, Y: 1, Gain: 22.0},
		{X: -3, Y: 1, Gain: 22.0},

		{X: -5, Y: 0, Gain: 22.0},
		{X: -4, Y: 0, Gain: 22.0},
		{X: -3, Y: 0, Gain: 22.0},
		{X: -2, Y: 0, Gain: 22.0},
		{X: -1, Y: 0, Gain: 22.0},

		{X: -4, Y: -1, Gain: 22.0},
		{X: -3, Y: -1, Gain: 22.0},

		{X: -3, Y: -2, Gain: 22.0},
	}
}

func (p *productionarrowleft) Category() beacons.Category {
	return beacons.Production
}

func (p *productionarrowleft) BType() beacons.BType {
	return beacons.Arrow
}

func (p *productionarrowleft) Image() image.Image { return img }

func init() {
	beacons.Add("l", func() beacons.Beacon { return &productionarrowleft{} })

	var err error
	img, err = beacons.FileToImage(assets, "data/ProductionArrowLeft.png")
	if err != nil {
		log.Fatalf("beacon image not found: %s", err.Error())
	}
}
