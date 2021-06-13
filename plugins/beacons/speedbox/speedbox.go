package speedbox

import (
	_ "embed"
	"image"
	"log"

	"github.com/MisterCodo/ngu/plugins/beacons"
)

//go:embed SpeedBox.png
var imageData []byte

type speedbox struct {
	effects []beacons.Effect
	img     image.Image
}

func (p *speedbox) Effect() []beacons.Effect {
	return p.effects
}

func (p *speedbox) Category() beacons.Category {
	return beacons.Speed
}

func (p *speedbox) BType() beacons.BType {
	return beacons.Box
}

func (p *speedbox) Image() image.Image { return p.img }

func init() {
	img, err := beacons.ImageFromBytes(imageData)
	if err != nil {
		log.Fatalf("beacon image error: %s", err.Error())
	}

	effects := []beacons.Effect{
		{X: -1, Y: 1, Gain: 40.0},
		{X: 0, Y: 1, Gain: 40.0},
		{X: 1, Y: 1, Gain: 40.0},
		{X: -1, Y: 0, Gain: 40.0},
		{X: 1, Y: 0, Gain: 40.0},
		{X: -1, Y: -1, Gain: 40.0},
		{X: 0, Y: -1, Gain: 40.0},
		{X: 1, Y: -1, Gain: 40.0},
	}

	beacons.Add("*", &speedbox{img: img, effects: effects})
}
