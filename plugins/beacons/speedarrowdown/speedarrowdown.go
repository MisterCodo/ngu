package speedarrowdown

import (
	"embed"
	"image"
	"log"

	"github.com/MisterCodo/ngu/plugins/beacons"
)

//go:embed data/*
var assets embed.FS

type speedarrowdown struct {
	effects []beacons.Effect
	img     image.Image
}

func (p *speedarrowdown) Effect() []beacons.Effect {
	return p.effects
}

func (p *speedarrowdown) Category() beacons.Category {
	return beacons.Speed
}

func (p *speedarrowdown) BType() beacons.BType {
	return beacons.Arrow
}

func (p *speedarrowdown) Image() image.Image { return p.img }

func init() {
	img, err := beacons.FileToImage(assets, "data/SpeedArrowDown.png")
	if err != nil {
		log.Fatalf("beacon image not found: %s", err.Error())
	}

	effects := []beacons.Effect{
		{X: -2, Y: 3, Gain: 26.0},

		{X: -1, Y: 4, Gain: 26.0},
		{X: -1, Y: 3, Gain: 26.0},

		{X: 0, Y: 5, Gain: 26.0},
		{X: 0, Y: 4, Gain: 26.0},
		{X: 0, Y: 3, Gain: 26.0},
		{X: 0, Y: 2, Gain: 26.0},
		{X: 0, Y: 1, Gain: 26.0},

		{X: 1, Y: 4, Gain: 26.0},
		{X: 1, Y: 3, Gain: 26.0},

		{X: 2, Y: 3, Gain: 26.0},
	}

	beacons.Add("v", &speedarrowdown{img: img, effects: effects})
}
