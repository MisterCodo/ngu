package productionarrowright

import (
	"github.com/MisterCodo/ngu/plugins/beacons"
)

type productionarrowright struct{}

func (p *productionarrowright) Effect() []beacons.Effect {
	return []beacons.Effect{
		{X: 3, Y: 2, Gain: 22.0},

		{X: 4, Y: 1, Gain: 22.0},
		{X: 3, Y: 1, Gain: 22.0},

		{X: 5, Y: 0, Gain: 22.0},
		{X: 4, Y: 0, Gain: 22.0},
		{X: 3, Y: 0, Gain: 22.0},
		{X: 2, Y: 0, Gain: 22.0},
		{X: 1, Y: 0, Gain: 22.0},

		{X: 4, Y: -1, Gain: 22.0},
		{X: 3, Y: -1, Gain: 22.0},

		{X: 3, Y: -2, Gain: 22.0},
	}
}

func (p *productionarrowright) Category() beacons.Category {
	return beacons.Production
}

func (p *productionarrowright) BType() beacons.BType {
	return beacons.Arrow
}

func init() {
	beacons.Add("r", func() beacons.Beacon { return &productionarrowright{} })
}
