package productionarrowleft

import (
	"github.com/MisterCodo/ngu/plugins/beacons"
)

type productionarrowleft struct{}

func (p *productionarrowleft) Effect() []beacons.Effect {
	return []beacons.Effect{
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
}

func init() {
	beacons.Add("l", func() beacons.Beacon { return &productionarrowleft{} })
}
