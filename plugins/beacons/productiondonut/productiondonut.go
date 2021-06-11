package productiondonut

import (
	"embed"
	"image"
	"log"

	"github.com/MisterCodo/ngu/plugins/beacons"
)

//go:embed data/*
var assets embed.FS
var img image.Image

type productiondonut struct{}

func (p *productiondonut) Effect() []beacons.Effect {
	return []beacons.Effect{
		{X: -2, Y: 2, Gain: 26.0},
		{X: -2, Y: 1, Gain: 26.0},
		{X: -2, Y: 0, Gain: 26.0},
		{X: -2, Y: -1, Gain: 26.0},
		{X: -2, Y: -2, Gain: 26.0},

		{X: -1, Y: 2, Gain: 26.0},
		{X: -1, Y: -2, Gain: 26.0},

		{X: 0, Y: 2, Gain: 26.0},
		{X: 0, Y: -2, Gain: 26.0},

		{X: 1, Y: 2, Gain: 26.0},
		{X: 1, Y: -2, Gain: 26.0},

		{X: 2, Y: 2, Gain: 26.0},
		{X: 2, Y: 1, Gain: 26.0},
		{X: 2, Y: 0, Gain: 26.0},
		{X: 2, Y: -1, Gain: 26.0},
		{X: 2, Y: -2, Gain: 26.0},
	}
}

func (p *productiondonut) Category() beacons.Category {
	return beacons.Production
}

func (p *productiondonut) BType() beacons.BType {
	return beacons.Donut
}

func (p *productiondonut) Image() image.Image { return img }

func init() {
	beacons.Add("O", func() beacons.Beacon { return &productiondonut{} })

	var err error
	img, err = beacons.FileToImage(assets, "data/ProductionDonut.png")
	if err != nil {
		log.Fatalf("beacon image not found: %s", err.Error())
	}
}
