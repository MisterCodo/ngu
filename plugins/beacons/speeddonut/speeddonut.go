package speeddonut

import (
	_ "embed"
	"image"
	"log"

	"github.com/MisterCodo/ngu/plugins/beacons"
)

//go:embed SpeedDonut.png
var imageData []byte

type speeddonut struct {
	effects []beacons.Effect
	img     image.Image
	name    string
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

func (p *speeddonut) Name() string { return p.name }

func init() {
	img, err := beacons.ImageFromBytes(imageData)
	if err != nil {
		log.Fatalf("beacon image error: %s", err.Error())
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

	beacons.Add("o", &speeddonut{img: img, effects: effects, name: "SpeedDonut"})
}
