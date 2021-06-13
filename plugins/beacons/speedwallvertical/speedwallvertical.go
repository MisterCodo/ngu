package speedwallvertical

import (
	"embed"
	"image"
	"log"

	"github.com/MisterCodo/ngu/plugins/beacons"
)

//go:embed data/*
var assets embed.FS

type speedwallvertical struct {
	effects []beacons.Effect
	img     image.Image
}

func (p *speedwallvertical) Effect() []beacons.Effect {
	return p.effects
}

func (p *speedwallvertical) Category() beacons.Category {
	return beacons.Speed
}

func (p *speedwallvertical) BType() beacons.BType {
	return beacons.Wall
}

func (p *speedwallvertical) Image() image.Image { return p.img }

func init() {
	img, err := beacons.FileToImage(assets, "data/SpeedWallVertical.png")
	if err != nil {
		log.Fatalf("beacon image not found: %s", err.Error())
	}

	effects := []beacons.Effect{
		{X: 0, Y: -6, Gain: 27.0},
		{X: 0, Y: -5, Gain: 27.0},
		{X: 0, Y: -4, Gain: 27.0},
		{X: 0, Y: -3, Gain: 27.0},
		{X: 0, Y: -2, Gain: 27.0},
		{X: 0, Y: -1, Gain: 27.0},

		{X: 0, Y: 1, Gain: 27.0},
		{X: 0, Y: 2, Gain: 27.0},
		{X: 0, Y: 3, Gain: 27.0},
		{X: 0, Y: 4, Gain: 27.0},
		{X: 0, Y: 5, Gain: 27.0},
		{X: 0, Y: 6, Gain: 27.0},
	}

	beacons.Add("|", &speedwallvertical{img: img, effects: effects})
}
