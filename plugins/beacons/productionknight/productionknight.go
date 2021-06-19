package productionknight

import (
	_ "embed"
	"image"
	"log"

	"github.com/MisterCodo/ngu/plugins/beacons"
)

//go:embed ProductionKnight.png
var imageData []byte

type productionknight struct {
	effects []beacons.Effect
	img     image.Image
	name    string
}

func (p *productionknight) Effect() []beacons.Effect {
	return p.effects
}

func (p *productionknight) Category() beacons.Category {
	return beacons.Production
}

func (p *productionknight) BType() beacons.BType {
	return beacons.Knight
}

func (p *productionknight) Image() image.Image { return p.img }

func (p *productionknight) Name() string { return p.name }

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

	beacons.Add("&", &productionknight{img: img, effects: effects, name: "ProductionKnight"})
}
