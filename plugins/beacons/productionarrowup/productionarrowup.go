package productionarrowup

import (
	_ "embed"
	"image"
	"log"

	"github.com/MisterCodo/ngu/plugins/beacons"
)

//go:embed ProductionArrowUp.png
var imageData []byte

type productionarrowup struct {
	effects []beacons.Effect
	img     image.Image
	name    string
}

func (p *productionarrowup) Effect() []beacons.Effect {
	return p.effects
}

func (p *productionarrowup) Category() beacons.Category {
	return beacons.Production
}

func (p *productionarrowup) BType() beacons.BType {
	return beacons.Arrow
}

func (p *productionarrowup) Image() image.Image { return p.img }

func (p *productionarrowup) Name() string { return p.name }

func init() {
	img, err := beacons.ImageFromBytes(imageData)
	if err != nil {
		log.Fatalf("beacon image error: %s", err.Error())
	}

	effects := []beacons.Effect{
		{X: -2, Y: -3, Gain: 22.0},

		{X: -1, Y: -4, Gain: 22.0},
		{X: -1, Y: -3, Gain: 22.0},

		{X: 0, Y: -5, Gain: 22.0},
		{X: 0, Y: -4, Gain: 22.0},
		{X: 0, Y: -3, Gain: 22.0},
		{X: 0, Y: -2, Gain: 22.0},
		{X: 0, Y: -1, Gain: 22.0},

		{X: 1, Y: -4, Gain: 22.0},
		{X: 1, Y: -3, Gain: 22.0},

		{X: 2, Y: -3, Gain: 22.0},
	}

	beacons.Add("u", &productionarrowup{img: img, effects: effects, name: "ProductionArrowUp"})
}
