package productionarrowright

import (
	_ "embed"
	"image"
	"log"

	"github.com/MisterCodo/ngu/plugins/beacons"
)

//go:embed ProductionArrowRight.png
var imageData []byte

type productionarrowright struct {
	effects []beacons.Effect
	img     image.Image
}

func (p *productionarrowright) Effect() []beacons.Effect {
	return p.effects
}

func (p *productionarrowright) Category() beacons.Category {
	return beacons.Production
}

func (p *productionarrowright) BType() beacons.BType {
	return beacons.Arrow
}

func (p *productionarrowright) Image() image.Image { return p.img }

func init() {
	img, err := beacons.ImageFromBytes(imageData)
	if err != nil {
		log.Fatalf("beacon image error: %s", err.Error())
	}

	effects := []beacons.Effect{
		{X: 3, Y: 2, Gain: 22.0},

		{X: 4, Y: 1, Gain: 22.0},
		{X: 3, Y: 1, Gain: 22.0},

		{X: 5, Y: 0, Gain: 22.0},
		{X: 4, Y: 0, Gain: 22.0},
		{X: 3, Y: 0, Gain: 22.0},
		{X: 2, Y: 0, Gain: 22.0},
		{X: 1, Y: 0, Gain: 22.0},

		{X: 4, Y: -1, Gain: 22.0},
		{X: 3, Y: -1, Gain: 22.0},

		{X: 3, Y: -2, Gain: 22.0},
	}

	beacons.Add("r", &productionarrowright{img: img, effects: effects})
}
