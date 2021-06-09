package beacons

type Beacon interface {
	Effect() []Effect
	Category() Category
	BType() BType
}

// Category represents the beacon category (Speed, Production or Efficiency)
type Category int

const (
	Speed = iota
	Production
	Efficiency
)

func (c Category) String() string {
	return [...]string{"Speed", "Production", "Efficiency"}[c]
}

// BType represents the beacon type (Box, Knight, Arrow, Wall, Donut)
type BType int

const (
	Box = iota
	Knight
	Arrow
	Wall
	Donut
)

func (bt BType) String() string {
	return [...]string{"Box", "Knight", "Arrow", "Wall", "Donut"}[bt]
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
