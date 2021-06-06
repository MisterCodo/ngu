package speedknight

import (
	"github.com/MisterCodo/ngu/plugins/beacons"
)

type speedknight struct{}

func (p *speedknight) Effect() []beacons.Effect {
	return []beacons.Effect{
		{X: -2, Y: 1, Gain: 35.0},
		{X: -1, Y: 2, Gain: 35.0},

		{X: 2, Y: 1, Gain: 35.0},
		{X: 1, Y: 2, Gain: 35.0},

		{X: -2, Y: -1, Gain: 35.0},
		{X: -1, Y: -2, Gain: 35.0},

		{X: 2, Y: -1, Gain: 35.0},
		{X: 1, Y: -2, Gain: 35.0},
	}
}

func init() {
	beacons.Add("k", func() beacons.Beacon { return &speedknight{} })
}
