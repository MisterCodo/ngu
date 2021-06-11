package productionarrowup

import (
	"embed"
	"image"
	"log"

	"github.com/MisterCodo/ngu/plugins/beacons"
)

//go:embed data/*
var assets embed.FS
var img image.Image

type productionarrowup struct{}

func (p *productionarrowup) Effect() []beacons.Effect {
	return []beacons.Effect{
		{X: -2, Y: -3, Gain: 22.0},

		{X: -1, Y: -4, Gain: 22.0},
		{X: -1, Y: -3, Gain: 22.0},

		{X: 0, Y: -5, Gain: 22.0},
		{X: 0, Y: -4, Gain: 22.0},
		{X: 0, Y: -3, Gain: 22.0},
		{X: 0, Y: -2, Gain: 22.0},
		{X: 0, Y: -1, Gain: 22.0},

		{X: 1, Y: -4, Gain: 22.0},
		{X: 1, Y: -3, Gain: 22.0},

		{X: 2, Y: -3, Gain: 22.0},
	}
}

func (p *productionarrowup) Category() beacons.Category {
	return beacons.Production
}

func (p *productionarrowup) BType() beacons.BType {
	return beacons.Arrow
}

func (p *productionarrowup) Image() image.Image { return img }

func init() {
	beacons.Add("u", func() beacons.Beacon { return &productionarrowup{} })

	var err error
	img, err = beacons.FileToImage(assets, "data/ProductionArrowUp.png")
	if err != nil {
		log.Fatalf("beacon image not found: %s", err.Error())
	}
}
