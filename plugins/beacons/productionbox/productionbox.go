package productionbox

import (
	"embed"
	"image"
	"log"

	"github.com/MisterCodo/ngu/plugins/beacons"
)

//go:embed data/*
var assets embed.FS

type productionbox struct {
	effects []beacons.Effect
	img     image.Image
}

func (p *productionbox) Effect() []beacons.Effect {
	return p.effects
}

func (p *productionbox) Category() beacons.Category {
	return beacons.Production
}

func (p *productionbox) BType() beacons.BType {
	return beacons.Box
}

func (p *productionbox) Image() image.Image { return p.img }

func init() {
	img, err := beacons.FileToImage(assets, "data/ProductionBox.png")
	if err != nil {
		log.Fatalf("beacon image not found: %s", err.Error())
	}

	effects := []beacons.Effect{
		{X: -1, Y: 1, Gain: 30.0},
		{X: 0, Y: 1, Gain: 30.0},
		{X: 1, Y: 1, Gain: 30.0},
		{X: -1, Y: 0, Gain: 30.0},
		{X: 1, Y: 0, Gain: 30.0},
		{X: -1, Y: -1, Gain: 30.0},
		{X: 0, Y: -1, Gain: 30.0},
		{X: 1, Y: -1, Gain: 30.0},
	}

	beacons.Add("b", &productionbox{img: img, effects: effects})
}
