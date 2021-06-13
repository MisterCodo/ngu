package speedarrowright

import (
	_ "embed"
	"image"
	"log"

	"github.com/MisterCodo/ngu/plugins/beacons"
)

//go:embed SpeedArrowRight.png
var imageData []byte

type speedarrowright struct {
	effects []beacons.Effect
	img     image.Image
}

func (p *speedarrowright) Effect() []beacons.Effect {
	return p.effects
}

func (p *speedarrowright) Category() beacons.Category {
	return beacons.Speed
}

func (p *speedarrowright) BType() beacons.BType {
	return beacons.Arrow
}

func (p *speedarrowright) Image() image.Image { return p.img }

func init() {
	img, err := beacons.ImageFromBytes(imageData)
	if err != nil {
		log.Fatalf("beacon image error: %s", err.Error())
	}

	effects := []beacons.Effect{
		{X: 3, Y: 2, Gain: 26.0},

		{X: 4, Y: 1, Gain: 26.0},
		{X: 3, Y: 1, Gain: 26.0},

		{X: 5, Y: 0, Gain: 26.0},
		{X: 4, Y: 0, Gain: 26.0},
		{X: 3, Y: 0, Gain: 26.0},
		{X: 2, Y: 0, Gain: 26.0},
		{X: 1, Y: 0, Gain: 26.0},

		{X: 4, Y: -1, Gain: 26.0},
		{X: 3, Y: -1, Gain: 26.0},

		{X: 3, Y: -2, Gain: 26.0},
	}

	beacons.Add(">", &speedarrowright{img: img, effects: effects})
}
