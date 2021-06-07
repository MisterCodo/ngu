package productionarrowup

import (
	"github.com/MisterCodo/ngu/plugins/beacons"
)

type productionarrowup struct{}

func (p *productionarrowup) Effect() []beacons.Effect {
	return []beacons.Effect{
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
}

func init() {
	beacons.Add("u", func() beacons.Beacon { return &productionarrowup{} })
}
