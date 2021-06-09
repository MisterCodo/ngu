package speedarrowdown

import (
	"github.com/MisterCodo/ngu/plugins/beacons"
)

type speedarrowdown struct{}

func (p *speedarrowdown) Effect() []beacons.Effect {
	return []beacons.Effect{
		{X: -2, Y: 3, Gain: 26.0},

		{X: -1, Y: 4, Gain: 26.0},
		{X: -1, Y: 3, Gain: 26.0},

		{X: 0, Y: 5, Gain: 26.0},
		{X: 0, Y: 4, Gain: 26.0},
		{X: 0, Y: 3, Gain: 26.0},
		{X: 0, Y: 2, Gain: 26.0},
		{X: 0, Y: 1, Gain: 26.0},

		{X: 1, Y: 4, Gain: 26.0},
		{X: 1, Y: 3, Gain: 26.0},

		{X: 2, Y: 3, Gain: 26.0},
	}
}

func (p *speedarrowdown) Category() beacons.Category {
	return beacons.Speed
}

func (p *speedarrowdown) BType() beacons.BType {
	return beacons.Arrow
}

func init() {
	beacons.Add("v", func() beacons.Beacon { return &speedarrowdown{} })
}
