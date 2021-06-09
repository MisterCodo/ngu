package productiondonut

import (
	"github.com/MisterCodo/ngu/plugins/beacons"
)

type productiondonut struct{}

func (p *productiondonut) Effect() []beacons.Effect {
	return []beacons.Effect{
		{X: -2, Y: 2, Gain: 26.0},
		{X: -2, Y: 1, Gain: 26.0},
		{X: -2, Y: 0, Gain: 26.0},
		{X: -2, Y: -1, Gain: 26.0},
		{X: -2, Y: -2, Gain: 26.0},

		{X: -1, Y: 2, Gain: 26.0},
		{X: -1, Y: -2, Gain: 26.0},

		{X: 0, Y: 2, Gain: 26.0},
		{X: 0, Y: -2, Gain: 26.0},

		{X: 1, Y: 2, Gain: 26.0},
		{X: 1, Y: -2, Gain: 26.0},

		{X: 2, Y: 2, Gain: 26.0},
		{X: 2, Y: 1, Gain: 26.0},
		{X: 2, Y: 0, Gain: 26.0},
		{X: 2, Y: -1, Gain: 26.0},
		{X: 2, Y: -2, Gain: 26.0},
	}
}

func (p *productiondonut) Category() beacons.Category {
	return beacons.Production
}

func (p *productiondonut) BType() beacons.BType {
	return beacons.Donut
}

func init() {
	beacons.Add("O", func() beacons.Beacon { return &productiondonut{} })
}
