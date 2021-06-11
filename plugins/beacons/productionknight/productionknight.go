package productionknight

import (
	"embed"
	"image"
	"log"

	"github.com/MisterCodo/ngu/plugins/beacons"
)

//go:embed data/*
var assets embed.FS
var img image.Image

type productionknight struct{}

func (p *productionknight) Effect() []beacons.Effect {
	return []beacons.Effect{
		{X: -2, Y: 1, Gain: 35.0},
		{X: -1, Y: 2, Gain: 35.0},

		{X: 2, Y: 1, Gain: 35.0},
		{X: 1, Y: 2, Gain: 35.0},

		{X: -2, Y: -1, Gain: 35.0},
		{X: -1, Y: -2, Gain: 35.0},

		{X: 2, Y: -1, Gain: 35.0},
		{X: 1, Y: -2, Gain: 35.0},
	}
}

func (p *productionknight) Category() beacons.Category {
	return beacons.Production
}

func (p *productionknight) BType() beacons.BType {
	return beacons.Knight
}

func (p *productionknight) Image() image.Image { return img }

func init() {
	beacons.Add("&", func() beacons.Beacon { return &productionknight{} })

	var err error
	img, err = beacons.FileToImage(assets, "data/ProductionKnight.png")
	if err != nil {
		log.Fatalf("beacon image not found: %s", err.Error())
	}
}
