package speedwallvertical

import (
	_ "embed"
	"image"
	"log"

	"github.com/MisterCodo/ngu/plugins/beacons"
)

//go:embed SpeedWallVertical.png
var imageData []byte

type speedwallvertical struct {
	effects []beacons.Effect
	img     image.Image
	name    string
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

func (p *speedwallvertical) Name() string { return p.name }

func init() {
	img, err := beacons.ImageFromBytes(imageData)
	if err != nil {
		log.Fatalf("beacon image error: %s", err.Error())
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

	beacons.Add("|", &speedwallvertical{img: img, effects: effects, name: "SpeedWallVertical"})
}
