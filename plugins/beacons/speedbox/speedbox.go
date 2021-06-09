package speedbox

import (
	"github.com/MisterCodo/ngu/plugins/beacons"
)

type speedbox struct{}

func (p *speedbox) Effect() []beacons.Effect {
	return []beacons.Effect{
		{X: -1, Y: 1, Gain: 40.0},
		{X: 0, Y: 1, Gain: 40.0},
		{X: 1, Y: 1, Gain: 40.0},
		{X: -1, Y: 0, Gain: 40.0},
		{X: 1, Y: 0, Gain: 40.0},
		{X: -1, Y: -1, Gain: 40.0},
		{X: 0, Y: -1, Gain: 40.0},
		{X: 1, Y: -1, Gain: 40.0},
	}
}

func (p *speedbox) Category() beacons.Category {
	return beacons.Speed
}

func (p *speedbox) BType() beacons.BType {
	return beacons.Box
}

func init() {
	beacons.Add("*", func() beacons.Beacon { return &speedbox{} })
}
