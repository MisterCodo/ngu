package speedarrowright

import (
	"github.com/MisterCodo/ngu/plugins/beacons"
)

type speedarrowright struct{}

func (p *speedarrowright) Effect() []beacons.Effect {
	return []beacons.Effect{
		{X: 3, Y: 2, Gain: 26.0},

		{X: 4, Y: 1, Gain: 26.0},
		{X: 3, Y: 1, Gain: 26.0},

		{X: 5, Y: 0, Gain: 26.0},
		{X: 4, Y: 0, Gain: 26.0},
		{X: 3, Y: 0, Gain: 26.0},
		{X: 2, Y: 0, Gain: 26.0},
		{X: 1, Y: 0, Gain: 26.0},

		{X: 4, Y: -1, Gain: 26.0},
		{X: 3, Y: -1, Gain: 26.0},

		{X: 3, Y: -2, Gain: 26.0},
	}
}

func init() {
	beacons.Add(">", func() beacons.Beacon { return &speedarrowright{} })
}
