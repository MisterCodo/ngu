package speedarrowup

import (
	_ "embed"
	"image"
	"log"

	"github.com/MisterCodo/ngu/plugins/beacons"
)

//go:embed SpeedArrowUp.png
var imageData []byte

type speedarrowup struct {
	effects []beacons.Effect
	img     image.Image
	name    string
}

func (p *speedarrowup) Effect() []beacons.Effect {
	return p.effects
}

func (p *speedarrowup) Category() beacons.Category {
	return beacons.Speed
}

func (p *speedarrowup) BType() beacons.BType {
	return beacons.Arrow
}

func (p *speedarrowup) Image() image.Image { return p.img }

func (p *speedarrowup) Name() string { return p.name }

func init() {
	img, err := beacons.ImageFromBytes(imageData)
	if err != nil {
		log.Fatalf("beacon image error: %s", err.Error())
	}

	effects := []beacons.Effect{
		{X: -2, Y: -3, Gain: 26.0},

		{X: -1, Y: -4, Gain: 26.0},
		{X: -1, Y: -3, Gain: 26.0},

		{X: 0, Y: -5, Gain: 26.0},
		{X: 0, Y: -4, Gain: 26.0},
		{X: 0, Y: -3, Gain: 26.0},
		{X: 0, Y: -2, Gain: 26.0},
		{X: 0, Y: -1, Gain: 26.0},

		{X: 1, Y: -4, Gain: 26.0},
		{X: 1, Y: -3, Gain: 26.0},

		{X: 2, Y: -3, Gain: 26.0},
	}

	beacons.Add("^", &speedarrowup{img: img, effects: effects, name: "SpeedArrowUp"})
}
