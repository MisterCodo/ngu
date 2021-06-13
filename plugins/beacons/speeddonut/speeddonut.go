package speeddonut

import (
	"embed"
	"image"
	"log"

	"github.com/MisterCodo/ngu/plugins/beacons"
)

//go:embed data/*
var assets embed.FS

type speeddonut struct {
	effects []beacons.Effect
	img     image.Image
}

func (p *speeddonut) Effect() []beacons.Effect {
	return p.effects
}

func (p *speeddonut) Category() beacons.Category {
	return beacons.Speed
}

func (p *speeddonut) BType() beacons.BType {
	return beacons.Donut
}

func (p *speeddonut) Image() image.Image { return p.img }

func init() {
	img, err := beacons.FileToImage(assets, "data/SpeedDonut.png")
	if err != nil {
		log.Fatalf("beacon image not found: %s", err.Error())
	}

	effects := []beacons.Effect{
		{X: -2, Y: 2, Gain: 23.0},
		{X: -2, Y: 1, Gain: 23.0},
		{X: -2, Y: 0, Gain: 23.0},
		{X: -2, Y: -1, Gain: 23.0},
		{X: -2, Y: -2, Gain: 23.0},

		{X: -1, Y: 2, Gain: 23.0},
		{X: -1, Y: -2, Gain: 23.0},

		{X: 0, Y: 2, Gain: 23.0},
		{X: 0, Y: -2, Gain: 23.0},

		{X: 1, Y: 2, Gain: 23.0},
		{X: 1, Y: -2, Gain: 23.0},

		{X: 2, Y: 2, Gain: 23.0},
		{X: 2, Y: 1, Gain: 23.0},
		{X: 2, Y: 0, Gain: 23.0},
		{X: 2, Y: -1, Gain: 23.0},
		{X: 2, Y: -2, Gain: 23.0},
	}

	beacons.Add("o", &speeddonut{img: img, effects: effects})
}
