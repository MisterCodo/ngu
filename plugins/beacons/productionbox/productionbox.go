package productionbox

import (
	"github.com/MisterCodo/ngu/plugins/beacons"
)

type productionbox struct{}

func (p *productionbox) Effect() []beacons.Effect {
	return []beacons.Effect{
		{X: -1, Y: 1, Gain: 30.0},
		{X: 0, Y: 1, Gain: 30.0},
		{X: 1, Y: 1, Gain: 30.0},
		{X: -1, Y: 0, Gain: 30.0},
		{X: 1, Y: 0, Gain: 30.0},
		{X: -1, Y: -1, Gain: 30.0},
		{X: 0, Y: -1, Gain: 30.0},
		{X: 1, Y: -1, Gain: 30.0},
	}
}

func (p *productionbox) Category() beacons.Category {
	return beacons.Production
}

func (p *productionbox) BType() beacons.BType {
	return beacons.Box
}

func init() {
	beacons.Add("b", func() beacons.Beacon { return &productionbox{} })
}
