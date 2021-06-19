package productiondonut

import (
	_ "embed"
	"image"
	"log"

	"github.com/MisterCodo/ngu/plugins/beacons"
)

//go:embed ProductionDonut.png
var imageData []byte

type productiondonut struct {
	effects []beacons.Effect
	img     image.Image
	name    string
}

func (p *productiondonut) Effect() []beacons.Effect {
	return p.effects
}

func (p *productiondonut) Category() beacons.Category {
	return beacons.Production
}

func (p *productiondonut) BType() beacons.BType {
	return beacons.Donut
}

func (p *productiondonut) Image() image.Image { return p.img }

func (p *productiondonut) Name() string { return p.name }

func init() {
	img, err := beacons.ImageFromBytes(imageData)
	if err != nil {
		log.Fatalf("beacon image error: %s", err.Error())
	}

	effects := []beacons.Effect{
		{X: -2, Y: 2, Gain: 26.0},
		{X: -2, Y: 1, Gain: 26.0},
		{X: -2, Y: 0, Gain: 26.0},
		{X: -2, Y: -1, Gain: 26.0},
		{X: -2, Y: -2, Gain: 26.0},

		{X: -1, Y: 2, Gain: 26.0},
		{X: -1, Y: -2, Gain: 26.0},

		{X: 0, Y: 2, Gain: 26.0},
		{X: 0, Y: -2, Gain: 26.0},

		{X: 1, Y: 2, Gain: 26.0},
		{X: 1, Y: -2, Gain: 26.0},

		{X: 2, Y: 2, Gain: 26.0},
		{X: 2, Y: 1, Gain: 26.0},
		{X: 2, Y: 0, Gain: 26.0},
		{X: 2, Y: -1, Gain: 26.0},
		{X: 2, Y: -2, Gain: 26.0},
	}

	beacons.Add("O", &productiondonut{img: img, effects: effects, name: "ProductionDonut"})
}
