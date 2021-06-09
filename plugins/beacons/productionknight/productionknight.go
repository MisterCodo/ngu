package productionknight

import (
	"github.com/MisterCodo/ngu/plugins/beacons"
)

type productionknight struct{}

func (p *productionknight) Effect() []beacons.Effect {
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

func (p *productionknight) Category() beacons.Category {
	return beacons.Production
}

func (p *productionknight) BType() beacons.BType {
	return beacons.Knight
}

func init() {
	beacons.Add("&", func() beacons.Beacon { return &productionknight{} })
}
