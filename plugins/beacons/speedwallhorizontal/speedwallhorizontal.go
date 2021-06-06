package speedwallhorizontal

import (
	"github.com/MisterCodo/ngu/plugins/beacons"
)

type speedwallhorizontal struct{}

func (p *speedwallhorizontal) Effect() []beacons.Effect {
	return []beacons.Effect{
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
}

func init() {
	beacons.Add("-", func() beacons.Beacon { return &speedwallhorizontal{} })
}
