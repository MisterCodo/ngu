package speeddonut

import (
	"github.com/MisterCodo/ngu/plugins/beacons"
)

type speeddonut struct{}

func (p *speeddonut) Effect() []beacons.Effect {
	return []beacons.Effect{
		{X: -2, Y: 2, Gain: 23.0},
		{X: -2, Y: 1, Gain: 23.0},
		{X: -2, Y: 0, Gain: 23.0},
		{X: -2, Y: -1, Gain: 23.0},
		{X: -2, Y: -2, Gain: 23.0},

		{X: -1, Y: 2, Gain: 23.0},
		{X: -1, Y: -2, Gain: 23.0},

		{X: 0, Y: 2, Gain: 23.0},
		{X: 0, Y: -2, Gain: 23.0},

		{X: 1, Y: 2, Gain: 23.0},
		{X: 1, Y: -2, Gain: 23.0},

		{X: 2, Y: 2, Gain: 23.0},
		{X: 2, Y: 1, Gain: 23.0},
		{X: 2, Y: 0, Gain: 23.0},
		{X: 2, Y: -1, Gain: 23.0},
		{X: 2, Y: -2, Gain: 23.0},
	}
}

func init() {
	beacons.Add("o", func() beacons.Beacon { return &speeddonut{} })
}
