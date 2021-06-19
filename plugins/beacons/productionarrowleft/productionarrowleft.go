package productionarrowleft

import (
	_ "embed"
	"image"
	"log"

	"github.com/MisterCodo/ngu/plugins/beacons"
)

//go:embed ProductionArrowLeft.png
var imageData []byte

type productionarrowleft struct {
	effects []beacons.Effect
	img     image.Image
	name    string
}

func (p *productionarrowleft) Effect() []beacons.Effect {
	return p.effects
}

func (p *productionarrowleft) Category() beacons.Category {
	return beacons.Production
}

func (p *productionarrowleft) BType() beacons.BType {
	return beacons.Arrow
}

func (p *productionarrowleft) Image() image.Image { return p.img }

func (p *productionarrowleft) Name() string { return p.name }

func init() {
	img, err := beacons.ImageFromBytes(imageData)
	if err != nil {
		log.Fatalf("beacon image error: %s", err.Error())
	}

	effects := []beacons.Effect{
		{X: -3, Y: 2, Gain: 22.0},

		{X: -4, Y: 1, Gain: 22.0},
		{X: -3, Y: 1, Gain: 22.0},

		{X: -5, Y: 0, Gain: 22.0},
		{X: -4, Y: 0, Gain: 22.0},
		{X: -3, Y: 0, Gain: 22.0},
		{X: -2, Y: 0, Gain: 22.0},
		{X: -1, Y: 0, Gain: 22.0},

		{X: -4, Y: -1, Gain: 22.0},
		{X: -3, Y: -1, Gain: 22.0},

		{X: -3, Y: -2, Gain: 22.0},
	}

	beacons.Add("l", &productionarrowleft{img: img, effects: effects, name: "ProductionArrowLeft"})
}
