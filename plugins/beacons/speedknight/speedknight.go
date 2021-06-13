package speedknight

import (
	"embed"
	"image"
	"log"

	"github.com/MisterCodo/ngu/plugins/beacons"
)

//go:embed data/*
var assets embed.FS

type speedknight struct {
	effects []beacons.Effect
	img     image.Image
}

func (p *speedknight) Effect() []beacons.Effect {
	return p.effects
}

func (p *speedknight) Category() beacons.Category {
	return beacons.Speed
}

func (p *speedknight) BType() beacons.BType {
	return beacons.Knight
}

func (p *speedknight) Image() image.Image { return p.img }

func init() {
	img, err := beacons.FileToImage(assets, "data/SpeedKnight.png")
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

	beacons.Add("k", &speedknight{img: img, effects: effects})
}
