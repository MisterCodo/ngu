package beacons

type Beacon interface {
	Effect() []Effect
}

// Genre represents the beacon type
type Genre int

const (
	Speed = iota
	Production
	Efficiency
)

func (g Genre) String() string {
	return [...]string{"Speed", "Production", "Efficiency"}[g]
}

// Effect indicates the impact at a location.
// X indicates moving left (X < 0) or right (X > 0)
// Y indicates moving down (Y < 0) or up (Y > 0)
// Gain indicates the gain percentage.
type Effect struct {
	X    int
	Y    int
	Gain float64
}
