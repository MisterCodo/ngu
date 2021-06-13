package productionknight

import (
	"embed"
	"image"
	"log"

	"github.com/MisterCodo/ngu/plugins/beacons"
)

//go:embed data/*
var assets embed.FS

type productionknight struct {
	effects []beacons.Effect
	img     image.Image
}

func (p *productionknight) Effect() []beacons.Effect {
	return p.effects
}

func (p *productionknight) Category() beacons.Category {
	return beacons.Production
}

func (p *productionknight) BType() beacons.BType {
	return beacons.Knight
}

func (p *productionknight) Image() image.Image { return p.img }

func init() {
	img, err := beacons.FileToImage(assets, "data/ProductionKnight.png")
	if err != nil {
		log.Fatalf("beacon image not found: %s", err.Error())
	}

	effects := []beacons.Effect{
		{X: -2, Y: 1, Gain: 35.0},
		{X: -1, Y: 2, Gain: 35.0},

		{X: 2, Y: 1, Gain: 35.0},
		{X: 1, Y: 2, Gain: 35.0},

		{X: -2, Y: -1, Gain: 35.0},
		{X: -1, Y: -2, Gain: 35.0},

		{X: 2, Y: -1, Gain: 35.0},
		{X: 1, Y: -2, Gain: 35.0},
	}

	beacons.Add("&", &productionknight{img: img, effects: effects})
}
