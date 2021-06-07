package productionwallvertical

import (
	"github.com/MisterCodo/ngu/plugins/beacons"
)

type productionwallvertical struct{}

func (p *productionwallvertical) Effect() []beacons.Effect {
	return []beacons.Effect{
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
	}
}

func init() {
	beacons.Add("w", func() beacons.Beacon { return &productionwallvertical{} })
}
