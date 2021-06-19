package productionwallhorizontal

import (
	_ "embed"
	"image"
	"log"

	"github.com/MisterCodo/ngu/plugins/beacons"
)

//go:embed ProductionWallHorizontal.png
var imageData []byte

type productionwallhorizontal struct {
	effects []beacons.Effect
	img     image.Image
	name    string
}

func (p *productionwallhorizontal) Effect() []beacons.Effect {
	return p.effects
}

func (p *productionwallhorizontal) Category() beacons.Category {
	return beacons.Production
}

func (p *productionwallhorizontal) BType() beacons.BType {
	return beacons.Wall
}

func (p *productionwallhorizontal) Image() image.Image { return p.img }

func (p *productionwallhorizontal) Name() string { return p.name }

func init() {
	img, err := beacons.ImageFromBytes(imageData)
	if err != nil {
		log.Fatalf("beacon image error: %s", err.Error())
	}

	effects := []beacons.Effect{
		{X: -6, Y: 0, Gain: 27.0},
		{X: -5, Y: 0, Gain: 27.0},
		{X: -4, Y: 0, Gain: 27.0},
		{X: -3, Y: 0, Gain: 27.0},
		{X: -2, Y: 0, Gain: 27.0},
		{X: -1, Y: 0, Gain: 27.0},

		{X: 1, Y: 0, Gain: 27.0},
		{X: 2, Y: 0, Gain: 27.0},
		{X: 3, Y: 0, Gain: 27.0},
		{X: 4, Y: 0, Gain: 27.0},
		{X: 5, Y: 0, Gain: 27.0},
		{X: 6, Y: 0, Gain: 27.0},
	}

	beacons.Add("h", &productionwallhorizontal{img: img, effects: effects, name: "SpeedWallHorizontal"})
}
