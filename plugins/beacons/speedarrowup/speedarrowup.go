package speedarrowup

import (
	"github.com/MisterCodo/ngu/plugins/beacons"
)

type speedarrowup struct{}

func (p *speedarrowup) Effect() []beacons.Effect {
	return []beacons.Effect{
		{X: -2, Y: -3, Gain: 26.0},

		{X: -1, Y: -4, Gain: 26.0},
		{X: -1, Y: -3, Gain: 26.0},

		{X: 0, Y: -5, Gain: 26.0},
		{X: 0, Y: -4, Gain: 26.0},
		{X: 0, Y: -3, Gain: 26.0},
		{X: 0, Y: -2, Gain: 26.0},
		{X: 0, Y: -1, Gain: 26.0},

		{X: 1, Y: -4, Gain: 26.0},
		{X: 1, Y: -3, Gain: 26.0},

		{X: 2, Y: -3, Gain: 26.0},
	}
}

func init() {
	beacons.Add("^", func() beacons.Beacon { return &speedarrowup{} })
}
