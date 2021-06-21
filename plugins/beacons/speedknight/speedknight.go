package speedknight

import (
	_ "embed"
	"image"
	"log"

	"github.com/MisterCodo/ngu/plugins/beacons"
)

//go:embed SpeedKnight.png
var imageData []byte

type speedknight struct {
	effects []beacons.Effect
	img     image.Image
	name    string
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

func (p *speedknight) Name() string { return p.name }

func init() {
	img, err := beacons.ImageFromBytes(imageData)
	if err != nil {
		log.Fatalf("beacon image error: %s", err.Error())
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

	beacons.Add("k", &speedknight{img: img, effects: effects, name: "SpeedKnight"})
}
