package speedarrowleft

import (
	_ "embed"
	"image"
	"log"

	"github.com/MisterCodo/ngu/plugins/beacons"
)

//go:embed SpeedArrowLeft.png
var imageData []byte

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
	img, err := beacons.ImageFromBytes(imageData)
	if err != nil {
		log.Fatalf("beacon image error: %s", err.Error())
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
