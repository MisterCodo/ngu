package productionwallvertical

import (
	_ "embed"
	"image"
	"log"

	"github.com/MisterCodo/ngu/plugins/beacons"
)

//go:embed ProductionWallVertical.png
var imageData []byte

type productionwallvertical struct {
	effects []beacons.Effect
	img     image.Image
	name    string
}

func (p *productionwallvertical) Effect() []beacons.Effect {
	return p.effects
}

func (p *productionwallvertical) Category() beacons.Category {
	return beacons.Production
}

func (p *productionwallvertical) BType() beacons.BType {
	return beacons.Wall
}

func (p *productionwallvertical) Image() image.Image { return p.img }

func (p *productionwallvertical) Name() string { return p.name }

func init() {
	img, err := beacons.ImageFromBytes(imageData)
	if err != nil {
		log.Fatalf("beacon image error: %s", err.Error())
	}

	effects := []beacons.Effect{
		{X: 0, Y: -6, Gain: 27.0},
		{X: 0, Y: -5, Gain: 27.0},
		{X: 0, Y: -4, Gain: 27.0},
		{X: 0, Y: -3, Gain: 27.0},
		{X: 0, Y: -2, Gain: 27.0},
		{X: 0, Y: -1, Gain: 27.0},

		{X: 0, Y: 1, Gain: 27.0},
		{X: 0, Y: 2, Gain: 27.0},
		{X: 0, Y: 3, Gain: 27.0},
		{X: 0, Y: 4, Gain: 27.0},
		{X: 0, Y: 5, Gain: 27.0},
		{X: 0, Y: 6, Gain: 27.0},
	}

	beacons.Add("w", &productionwallvertical{img: img, effects: effects, name: "ProductionWallVertical"})
}
