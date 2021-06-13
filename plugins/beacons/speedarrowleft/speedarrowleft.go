package speedarrowleft

import (
	"embed"
	"image"
	"log"

	"github.com/MisterCodo/ngu/plugins/beacons"
)

//go:embed data/*
var assets embed.FS

type speedarrowleft struct {
	effects []beacons.Effect
	img     image.Image
}

func (p *speedarrowleft) Effect() []beacons.Effect {
	return p.effects
}

func (p *speedarrowleft) Category() beacons.Category {
	return beacons.Speed
}

func (p *speedarrowleft) BType() beacons.BType {
	return beacons.Arrow
}

func (p *speedarrowleft) Image() image.Image { return p.img }

func init() {
	img, err := beacons.FileToImage(assets, "data/SpeedArrowLeft.png")
	if err != nil {
		log.Fatalf("beacon image not found: %s", err.Error())
	}

	effects := []beacons.Effect{
		{X: -3, Y: 2, Gain: 26.0},

		{X: -4, Y: 1, Gain: 26.0},
		{X: -3, Y: 1, Gain: 26.0},

		{X: -5, Y: 0, Gain: 26.0},
		{X: -4, Y: 0, Gain: 26.0},
		{X: -3, Y: 0, Gain: 26.0},
		{X: -2, Y: 0, Gain: 26.0},
		{X: -1, Y: 0, Gain: 26.0},

		{X: -4, Y: -1, Gain: 26.0},
		{X: -3, Y: -1, Gain: 26.0},

		{X: -3, Y: -2, Gain: 26.0},
	}

	beacons.Add("<", &speedarrowleft{img: img, effects: effects})
}
