package speedwallhorizontal

import (
	_ "embed"
	"image"
	"log"

	"github.com/MisterCodo/ngu/plugins/beacons"
)

//go:embed SpeedWallHorizontal.png
var imageData []byte

type speedwallhorizontal struct {
	effects []beacons.Effect
	img     image.Image
	name    string
}

func (p *speedwallhorizontal) Effect() []beacons.Effect {
	return p.effects
}

func (p *speedwallhorizontal) Category() beacons.Category {
	return beacons.Speed
}

func (p *speedwallhorizontal) BType() beacons.BType {
	return beacons.Wall
}

func (p *speedwallhorizontal) Image() image.Image { return p.img }

func (p *speedwallhorizontal) Name() string { return p.name }

func init() {
	img, err := beacons.ImageFromBytes(imageData)
	if err != nil {
		log.Fatalf("beacon image error: %s", err.Error())
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

	beacons.Add("-", &speedwallhorizontal{img: img, effects: effects, name: "SpeedWallHorizontal"})
}
