package productionwallhorizontal

import (
	"github.com/MisterCodo/ngu/plugins/beacons"
)

type productionwallhorizontal struct{}

func (p *productionwallhorizontal) Effect() []beacons.Effect {
	return []beacons.Effect{
		{X: -6, Y: 0, Gain: 27.0},
		{X: -5, Y: 0, Gain: 27.0},
		{X: -4, Y: 0, Gain: 27.0},
		{X: -3, Y: 0, Gain: 27.0},
		{X: -2, Y: 0, Gain: 27.0},
		{X: -1, Y: 0, Gain: 27.0},

		{X: 1, Y: 0, Gain: 27.0},
		{X: 2, Y: 0, Gain: 27.0},
		{X: 3, Y: 0, Gain: 27.0},
		{X: 4, Y: 0, Gain: 27.0},
		{X: 5, Y: 0, Gain: 27.0},
		{X: 6, Y: 0, Gain: 27.0},
	}
}

func (p *productionwallhorizontal) Category() beacons.Category {
	return beacons.Production
}

func (p *productionwallhorizontal) BType() beacons.BType {
	return beacons.Wall
}

func init() {
	beacons.Add("h", func() beacons.Beacon { return &productionwallhorizontal{} })
}
