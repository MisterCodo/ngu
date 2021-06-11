package productionarrowright

import (
	"embed"
	"image"
	"log"

	"github.com/MisterCodo/ngu/plugins/beacons"
)

//go:embed data/*
var assets embed.FS
var img image.Image

type productionarrowright struct{}

func (p *productionarrowright) Effect() []beacons.Effect {
	return []beacons.Effect{
		{X: 3, Y: 2, Gain: 22.0},

		{X: 4, Y: 1, Gain: 22.0},
		{X: 3, Y: 1, Gain: 22.0},

		{X: 5, Y: 0, Gain: 22.0},
		{X: 4, Y: 0, Gain: 22.0},
		{X: 3, Y: 0, Gain: 22.0},
		{X: 2, Y: 0, Gain: 22.0},
		{X: 1, Y: 0, Gain: 22.0},

		{X: 4, Y: -1, Gain: 22.0},
		{X: 3, Y: -1, Gain: 22.0},

		{X: 3, Y: -2, Gain: 22.0},
	}
}

func (p *productionarrowright) Category() beacons.Category {
	return beacons.Production
}

func (p *productionarrowright) BType() beacons.BType {
	return beacons.Arrow
}

func (p *productionarrowright) Image() image.Image { return img }

func init() {
	beacons.Add("r", func() beacons.Beacon { return &productionarrowright{} })

	var err error
	img, err = beacons.FileToImage(assets, "data/ProductionArrowRight.png")
	if err != nil {
		log.Fatalf("beacon image not found: %s", err.Error())
	}
}
