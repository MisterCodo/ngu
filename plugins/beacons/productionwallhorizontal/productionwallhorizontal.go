package productionwallhorizontal

import (
	"embed"
	"image"
	"log"

	"github.com/MisterCodo/ngu/plugins/beacons"
)

//go:embed data/*
var assets embed.FS

type productionwallhorizontal struct {
	effects []beacons.Effect
	img     image.Image
}

func (p *productionwallhorizontal) Effect() []beacons.Effect {
	return p.effects
}

func (p *productionwallhorizontal) Category() beacons.Category {
	return beacons.Production
}

func (p *productionwallhorizontal) BType() beacons.BType {
	return beacons.Wall
}

func (p *productionwallhorizontal) Image() image.Image { return p.img }

func init() {
	img, err := beacons.FileToImage(assets, "data/ProductionWallHorizontal.png")
	if err != nil {
		log.Fatalf("beacon image not found: %s", err.Error())
	}

	effects := []beacons.Effect{
		{X: -6, Y: 0, Gain: 27.0},
		{X: -5, Y: 0, Gain: 27.0},
		{X: -4, Y: 0, Gain: 27.0},
		{X: -3, Y: 0, Gain: 27.0},
		{X: -2, Y: 0, Gain: 27.0},
		{X: -1, Y: 0, Gain: 27.0},

		{X: 1, Y: 0, Gain: 27.0},
		{X: 2, Y: 0, Gain: 27.0},
		{X: 3, Y: 0, Gain: 27.0},
		{X: 4, Y: 0, Gain: 27.0},
		{X: 5, Y: 0, Gain: 27.0},
		{X: 6, Y: 0, Gain: 27.0},
	}

	beacons.Add("h", &productionwallhorizontal{img: img, effects: effects})
}
