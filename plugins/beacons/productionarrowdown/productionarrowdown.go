package productionarrowdown

import (
	"github.com/MisterCodo/ngu/plugins/beacons"
)

type productionarrowdown struct{}

func (p *productionarrowdown) Effect() []beacons.Effect {
	return []beacons.Effect{
		{X: -2, Y: 3, Gain: 22.0},

		{X: -1, Y: 4, Gain: 22.0},
		{X: -1, Y: 3, Gain: 22.0},

		{X: 0, Y: 5, Gain: 22.0},
		{X: 0, Y: 4, Gain: 22.0},
		{X: 0, Y: 3, Gain: 22.0},
		{X: 0, Y: 2, Gain: 22.0},
		{X: 0, Y: 1, Gain: 22.0},

		{X: 1, Y: 4, Gain: 22.0},
		{X: 1, Y: 3, Gain: 22.0},

		{X: 2, Y: 3, Gain: 22.0},
	}
}

func (p *productionarrowdown) Category() beacons.Category {
	return beacons.Production
}

func (p *productionarrowdown) BType() beacons.BType {
	return beacons.Arrow
}

func init() {
	beacons.Add("d", func() beacons.Beacon { return &productionarrowdown{} })
}
