package productionbox

import (
	_ "embed"
	"image"
	"log"

	"github.com/MisterCodo/ngu/plugins/beacons"
)

//go:embed ProductionBox.png
var imageData []byte

type productionbox struct {
	effects []beacons.Effect
	img     image.Image
	name    string
}

func (p *productionbox) Effect() []beacons.Effect {
	return p.effects
}

func (p *productionbox) Category() beacons.Category {
	return beacons.Production
}

func (p *productionbox) BType() beacons.BType {
	return beacons.Box
}

func (p *productionbox) Image() image.Image { return p.img }

func (p *productionbox) Name() string { return p.name }

func init() {
	img, err := beacons.ImageFromBytes(imageData)
	if err != nil {
		log.Fatalf("beacon image error: %s", err.Error())
	}

	effects := []beacons.Effect{
		{X: -1, Y: 1, Gain: 30.0},
		{X: 0, Y: 1, Gain: 30.0},
		{X: 1, Y: 1, Gain: 30.0},
		{X: -1, Y: 0, Gain: 30.0},
		{X: 1, Y: 0, Gain: 30.0},
		{X: -1, Y: -1, Gain: 30.0},
		{X: 0, Y: -1, Gain: 30.0},
		{X: 1, Y: -1, Gain: 30.0},
	}

	beacons.Add("b", &productionbox{img: img, effects: effects, name: "ProductionBox"})
}
