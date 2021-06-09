package speedwallvertical

import (
	"github.com/MisterCodo/ngu/plugins/beacons"
)

type speedwallvertical struct{}

func (p *speedwallvertical) Effect() []beacons.Effect {
	return []beacons.Effect{
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
}

func (p *speedwallvertical) Category() beacons.Category {
	return beacons.Speed
}

func (p *speedwallvertical) BType() beacons.BType {
	return beacons.Wall
}

func init() {
	beacons.Add("|", func() beacons.Beacon { return &speedwallvertical{} })
}
