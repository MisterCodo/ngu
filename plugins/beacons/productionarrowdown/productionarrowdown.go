package productionarrowdown

import (
	"embed"
	"image"
	"log"

	"github.com/MisterCodo/ngu/plugins/beacons"
)

//go:embed data/*
var assets embed.FS

type productionarrowdown struct {
	effects []beacons.Effect
	img     image.Image
}

func (p *productionarrowdown) Effect() []beacons.Effect {
	return p.effects
}

func (p *productionarrowdown) Category() beacons.Category {
	return beacons.Production
}

func (p *productionarrowdown) BType() beacons.BType {
	return beacons.Arrow
}

func (p *productionarrowdown) Image() image.Image { return p.img }

func init() {
	img, err := beacons.FileToImage(assets, "data/ProductionArrowDown.png")
	if err != nil {
		log.Fatalf("beacon image not found: %s", err.Error())
	}

	effects := []beacons.Effect{
		{X: -2, Y: 3, Gain: 22.0},

		{X: -1, Y: 4, Gain: 22.0},
		{X: -1, Y: 3, Gain: 22.0},

		{X: 0, Y: 5, Gain: 22.0},
		{X: 0, Y: 4, Gain: 22.0},
		{X: 0, Y: 3, Gain: 22.0},
		{X: 0, Y: 2, Gain: 22.0},
		{X: 0, Y: 1, Gain: 22.0},

		{X: 1, Y: 4, Gain: 22.0},
		{X: 1, Y: 3, Gain: 22.0},

		{X: 2, Y: 3, Gain: 22.0},
	}

	beacons.Add("d", &productionarrowdown{img: img, effects: effects})
}
